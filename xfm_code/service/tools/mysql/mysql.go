package dmrdb

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/patrickmn/go-cache"
)

// 封装同步,如用于规则缓存是为了多协程都判断不在在时,防止多次刷新同一份元数据
type syncCache struct {
	Cache *cache.Cache
	MU    sync.Mutex
}

// 默认不设过期时间
func New() *syncCache {
	return &syncCache{Cache: cache.New(cache.NoExpiration, cache.NoExpiration)}
}

// 默认不设置过期时间
func (sc *syncCache) Set(k string, x interface{}) {
	sc.Cache.Set(k, x, cache.NoExpiration)
}

func (sc *syncCache) Get(k string) (interface{}, bool) {
	return sc.Cache.Get(k)
}

//关系型数据库通用包

const INDEX_MYSQL = "mysql"
const INDEX_POSTGRES = "postgres"

type RDBCli struct {
	*sql.DB
}

// map[string]interface{}
var rDBCliPool = New()

// 得到Mysql数据库的连接,传入连接串及最大连接数
func GetMysqlPoolCli(connUrl string, maxOpen int) *RDBCli {
	return getConnPoolByrDBInfo(INDEX_MYSQL, connUrl, maxOpen)
}

// 得到TPG数据库的连接,传入连接串及最大连接数
func GetPostgresPoolCli(connUrl string, maxOpen int) *RDBCli {
	return getConnPoolByrDBInfo(INDEX_POSTGRES, connUrl, maxOpen)
}

/*
返回connUrl对应的连接池
rdbType:关系型数据库类型=mysql|postgres
connUrl:连接串
maxOpen:连接池最大连接数
例子:
mysql=Datamore:datamore_oss@tcp(proxy.tencent-cloud.net:25006)/db_cf_logdb?charset=utf8
postgres=user=u_hy_dc_oss password=u_hy_dc_oss_da dbname=tdwdata sslmode=disable host=tdwdata-bi-tdw.oa.com
*/
func getConnPoolByrDBInfo(rDBType, connUrl string, maxOpen int) *RDBCli {
	key := connUrl
	rdbCli, ok := rDBCliPool.Get(key)
	if ok && rdbCli != nil {
		return rdbCli.(*RDBCli)
	}
	rDBCliPool.MU.Lock()
	defer rDBCliPool.MU.Unlock()

	rdbCli, ok = rDBCliPool.Get(key)
	if ok && rdbCli != nil {
		return rdbCli.(*RDBCli)
	}
	rdbCliNew := generateRDBPoolByInfo(rDBType, connUrl, maxOpen)
	rDBCliPool.Set(key, rdbCliNew)
	return rdbCliNew
}

func generateRDBPoolByInfo(rDBType, connUrl string, maxOpen int) (rDBCli *RDBCli) {
	log.Println("generateRDBPoolByInfo,rDBType:", rDBType, ",connUrl:", connUrl, ",maxOpen:", maxOpen)
	db, dbErr := sql.Open(rDBType, connUrl)
	if dbErr != nil {
		log.Println("generateRDBPoolByInfo_dbErr:", dbErr, ",connUrl:", connUrl)
		if db != nil {
			db.Close()
		}
		panic(dbErr)
	}
	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxOpen)
	db.SetConnMaxLifetime(time.Minute * 30)
	rDBCli = &RDBCli{db}
	return rDBCli
}

// 通用方法,传入sql直接查询返回map切片
func (rDBCli *RDBCli) QuerySqlForMap(sqlString string, queryTimeout time.Duration, args ...interface{}) ([]map[string]string, error) {
	resultMap := make([]map[string]string, 0, 10)
	ctx, _ := context.WithTimeout(context.Background(), queryTimeout)
	rows, queryErr := rDBCli.QueryContext(ctx, sqlString, args...)
	if queryErr != nil {
		log.Println("QuerySqlForMap_queryErr:", queryErr, ",sql:", sqlString)
		return resultMap, queryErr
	}
	defer rows.Close()
	rowCols, _ := rows.Columns()

	values := make([]sql.RawBytes, len(rowCols))
	scans := make([]interface{}, len(rowCols))
	for i := range values {
		scans[i] = &values[i]
	}

	for rows.Next() {
		rowErr := rows.Scan(scans...)
		if rowErr != nil {
			log.Println("QuerySqlForMap_rowErr:", rowErr, ",sql:", sqlString)
			return resultMap, rowErr
		}
		each := make(map[string]string, len(scans))
		for i, col := range values {
			each[strings.ToLower(rowCols[i])] = string(col)
		}

		resultMap = append(resultMap, each)
	}
	lastErr := rows.Err()
	if lastErr != nil {
		log.Println("QuerySqlForMap_lastErr:", lastErr, ",sql:", sqlString)
		return resultMap, lastErr
	}
	return resultMap, nil
}

// 通用方法,传入sql直接查询返回map切片
func (rDBCli *RDBCli) QuerySqlForOrigalMap(sqlString string, queryTimeout time.Duration, args ...interface{}) ([]map[string]interface{}, error) {
	resultMap := make([]map[string]interface{}, 0, 10)
	ctx, _ := context.WithTimeout(context.Background(), queryTimeout)
	stmt, queryErr := rDBCli.PrepareContext(ctx, sqlString)
	defer stmt.Close()
	if queryErr != nil {
		log.Println("QuerySqlForOrigalMap_queryErr:", queryErr, ",sql:", sqlString)
		return resultMap, queryErr
	}
	rows, rowsErr := stmt.Query(args...)
	defer rows.Close()
	if rowsErr != nil {
		log.Println("QuerySqlForOrigalMap_rowsErr:", rowsErr, ",sql:", sqlString)
		return resultMap, rowsErr
	}
	rowCols, _ := rows.Columns()

	values := make([]interface{}, len(rowCols))
	scans := make([]interface{}, len(rowCols))
	for i := range values {
		scans[i] = &values[i]
	}

	for rows.Next() {
		rowErr := rows.Scan(scans...)
		if rowErr != nil {
			log.Println("QuerySqlForMap_rowErr:", rowErr, ",sql:", sqlString)
			return resultMap, rowErr
		}
		each := make(map[string]interface{}, len(scans))
		for i, col := range values {
			b, ok := col.([]byte)
			if ok {
				each[strings.ToLower(rowCols[i])] = string(b)
			} else {
				each[strings.ToLower(rowCols[i])] = col
			}
		}
		resultMap = append(resultMap, each)
	}
	lastErr := rows.Err()
	if lastErr != nil {
		log.Println("QuerySqlForMap_lastErr:", lastErr, ",sql:", sqlString)
		return resultMap, lastErr
	}
	return resultMap, nil
}

// QuerySqlForSingleString 适用于只返回一个值的sql,如sum,count这一类等等,当sql返回的是null值是,对应返回的是""空字符串
func (rDBCli *RDBCli) QuerySqlForSingleString(sqlString string, queryTimeout time.Duration, args ...interface{}) (result string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), queryTimeout)
	rows, queryErr := rDBCli.QueryContext(ctx, sqlString, args...)
	if queryErr != nil {
		log.Println("QuerySqlForSingleString_queryErr:", queryErr, ",sql:", sqlString)
		return result, queryErr
	}
	defer rows.Close()
	var value sql.RawBytes
	if rows.Next() { //只取第一个值即可
		rowErr := rows.Scan(&value)
		if rowErr != nil {
			log.Println("QuerySqlForSingleString_rowErr:", rowErr, ",sql:", sqlString)
			return result, rowErr
		}
	}
	lastErr := rows.Err()
	if lastErr != nil {
		log.Println("QuerySqlForSingleString_lastErr:", lastErr, ",sql:", sqlString)
		return result, lastErr
	}
	result = string(value)
	return result, nil
}

// 执行数据库更新操作
func (rDBCli *RDBCli) ExecuteUpdate(updateSql string, updateTimeout time.Duration, args ...interface{}) (int64, error) {
	ctx, _ := context.WithTimeout(context.Background(), updateTimeout)
	res, resErr := rDBCli.ExecContext(ctx, updateSql, args...)
	if resErr != nil {
		log.Println("ExecuteUpdate_resErr:", resErr, ",updateSql:", updateSql, ",updateTimeout:", updateTimeout)
		return 0, resErr
	}
	return res.RowsAffected()
}

// 获取缓存的值,管理用途
func GetConnPoolCache(connUrl string) (showMap map[string]*RDBCli) {
	rDBCliPool.MU.Lock()
	defer rDBCliPool.MU.Unlock()

	showMap = make(map[string]*RDBCli, rDBCliPool.Cache.ItemCount())

	for k, v := range rDBCliPool.Cache.Items() {
		if connUrl == "0" || strings.Index(k, connUrl) >= 0 {
			showMap[k] = v.Object.(*RDBCli)
		}
	}
	return showMap
}
