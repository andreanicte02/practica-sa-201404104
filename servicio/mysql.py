import pymysql
import pymysql.cursors


class Database:
    def __init__(self):
        self.connection = pymysql.connect(
            host='database',
            user="root",
            passwd="201404104",
            db="p8sa",
            cursorclass=pymysql.cursors.DictCursor,
            sql_mode=''
        )

    def get_simple(self, reporte):
        ret = {}
        try:
            with self.connection.cursor() as cursor:
                sql = 'SELECT *From Estudiante'
                cursor.execute(sql)
                ret = cursor.fetchall()


        except Exception as e:
            ret = {}
            ret['ok'] = False
            ret['error'] = str(e)
        return ret