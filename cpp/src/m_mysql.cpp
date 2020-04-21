#include "m_mysql.h"

int M_Mysql::openConnPool(int coreConnNum)
{
    if (coreConnNum > MAX_POOL_SIZE || coreConnNum <= 0)
    {
        coreConnNum = MAX_POOL_SIZE;
    }
    pthread_mutex_init(&mutex, NULL);

    connNum = coreConnNum;
    sem_init(&sem, 0, connNum);

    for (int i = 0; i < connNum; i++)
    {
        mysqlConn *conn = new mysqlConn;
        conn->connSetting = connSetting;
        if (mysql_init(&conn->mysql) == NULL)
        {
            MCP_LOG("ERROR: mysql_init() %s\n", mysql_error(&conn->mysql));
            return 1;
        }
        mysql_options(&conn->mysql, MYSQL_OPT_CONNECT_TIMEOUT, &connSetting->timeout);
        mysql_options(&conn->mysql, MYSQL_OPT_READ_TIMEOUT, &connSetting->timeout);
        mysql_options(&conn->mysql, MYSQL_SET_CHARSET_NAME, connSetting->charset);
        conn->sock = mysql_real_connect(&conn->mysql, connSetting->host, connSetting->user, connSetting->password, connSetting->database, connSetting->port, NULL, CLIENT_MULTI_STATEMENTS);
        if (!conn->sock)
        {
            MCP_LOG("ERROR mysql_real_connect(): %s\n", mysql_error(&conn->mysql));
            return 1;
        }
        connPool.push_back(conn);
    }
    return 0;
}

void M_Mysql::closeConnPool()
{
    pthread_mutex_destroy(&mutex);
    std::deque<mysqlConn *>::iterator iter;
    for (iter = connPool.begin(); iter != connPool.end(); ++iter)
    {
        mysqlConn *conn = *iter;
        if (conn->sock != NULL)
        {
            mysql_close(conn->sock);
            conn->sock = NULL;
        }
        delete conn;
    }
    connPool.clear();
    sem_destroy(&sem);
    free(connSetting);
}

int M_Mysql::lockPool()
{
    struct timespec ts;
    if (clock_gettime(CLOCK_REALTIME, &ts) == -1)
    {
        MCP_LOG("Function clock_gettime failed");
        return -1;
    }

    ts.tv_sec += 1;
    int ret = 0;
    while ((ret = sem_timedwait(&sem, &ts)) == -1 && errno == EINTR)
        continue; // Restart when interrupted by handler
    if (ret == -1)
    {
        if (errno == ETIMEDOUT)
        {
            MCP_LOG("Timeout occurred in locking connection, \
                database is, pool connNum is %d",
                    connNum);
        }
        else
        {
            MCP_LOG("Unknown error in locking connection, \
                database is, pool connNum is %d",
                    connNum);
        }
        return -2;
    }
    return 0;
}

mysqlConn *M_Mysql::fetchConnection()
{
    if (lockPool() != 0)
    {
        return NULL;
    }
    pthread_mutex_lock(&mutex);

    mysqlConn *conn = connPool.front();
    connPool.pop_front();

    pthread_mutex_unlock(&mutex);
    return conn;
}

int M_Mysql::executeSql(mysqlConn *conn, const char *sql)
{
    if (NULL == conn || NULL == sql)
    {
        return 1;
    }
    if (conn->sock)
    {
        conn->res = mysql_query(conn->sock, sql);
    }
    else
    {
        conn->res = 1;
    }
    //reconnect
    if (conn->res)
    {
        mysql_close(conn->sock);
        mysql_init(&(conn->mysql));
        setting *s = conn->connSetting;
        mysql_options(&conn->mysql, MYSQL_OPT_CONNECT_TIMEOUT, &s->timeout);
        mysql_options(&conn->mysql, MYSQL_OPT_READ_TIMEOUT, &s->timeout);
        mysql_options(&conn->mysql, MYSQL_SET_CHARSET_NAME, s->charset);
        conn->sock = mysql_real_connect(&(conn->mysql), s->host, s->user, s->password, s->database, s->port, NULL, CLIENT_MULTI_STATEMENTS);
        if (!conn->sock)
        {
            MCP_LOG("Failed to connect to database: Error: %s, database is %s\n", mysql_error(&conn->mysql), conn->connSetting->database);
        }
        conn->res = mysql_query(conn->sock, sql);
    }
    return conn->res;
}

//recycle connection
void M_Mysql::recycleConnection(mysqlConn *conn)
{
    pthread_mutex_lock(&mutex);
    connPool.push_back(conn);
    pthread_mutex_unlock(&mutex);
    sem_post(&sem);
    return;
}

//初始化mysql
int M_Mysql::initMysqlConnPool(const char *host, int port, const char *user, const char *password, const char *database)
{
    if (NULL == host)
    {
        MCP_LOG("ERROR: host is NULL");
        return INIT_ERROR;
    }
    if (NULL == user)
    {
        MCP_LOG("ERROR: user is NULL");
        return INIT_ERROR;
    }
    if (NULL == password)
    {
        MCP_LOG("ERROR: password is NULL");
        return INIT_ERROR;
    }
    if (NULL == database)
    {
        MCP_LOG("ERROR: database is NULL");
        return INIT_ERROR;
    }
    connSetting = (setting *)malloc(sizeof(setting));
    assert(connSetting != NULL);
    memset(connSetting, 0, sizeof(setting));
    strncpy(connSetting->host, host, MAX_SETTING_STRING_LEN);
    strncpy(connSetting->user, user, MAX_SETTING_STRING_LEN);
    strncpy(connSetting->password, password, MAX_SETTING_STRING_LEN);
    connSetting->port = port;
    strncpy(connSetting->database, database, MAX_SETTING_STRING_LEN);
    sprintf(connSetting->charset, "utf8");
    connSetting->timeout = 2;
    return INIT_SUCCESS;
}

//set the connection charset
void M_Mysql::setCharsetOption(setting *cs, const char *charset)
{
    if (NULL != cs && NULL != charset)
    {
        strncpy(cs->charset, charset, MAX_SETTING_STRING_LEN);
    }
}