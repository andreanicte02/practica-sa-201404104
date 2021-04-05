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

    def add_reporte(self, data):
        ret = {}
        try:
            with self.connection.cursor() as cursor:
                sql = '''INSERT INTO Estudiante(carnet, nombre, curso)
                    SELECT * FROM (SELECT %s,%s,%s)AS tmp
                    WHERE NOT EXISTS(
                        SELECT carnet FROM Estudiante WHERE carnet = %s
                    ) LIMIT 1'''
                cursor.execute(
                    sql, (data['carnet'], data['nombre'], data['curso'], data['carnet']))

                sql2 = '''SELECT * FROM Estudiante WHERE carnet = %s'''
                cursor.execute(sql2, (data['carnet'],))
                ret = cursor.fetchone()

                sql3 = '''INSERT INTO Reporte(mensaje, procesado, fecha, estudiante) VALUES(%s,%s,CURDATE(),%s)'''
                cursor.execute(
                    sql3, (data['mensaje'], data['procesado'], ret['ID_estudiante']))

            self.connection.commit()
            ret['ok'] = True
        except Exception as e:
            ret = {}
            ret['ok'] = False
            ret['error'] = str(e)
        return ret

    def get_reporte(self, reporte):
        ret = {}
        try:
            with self.connection.cursor() as cursor:
                sql = 'SELECT e.carnet, e.nombre, e.curso, r.procesado, r.fecha, r.mensaje FROM Estudiante e INNER JOIN Reporte r ON r.estudiante = e.ID_estudiante WHERE r.ID_reporte = %s'
                cursor.execute(sql, (reporte,))
                ret = cursor.fetchone()
        except Exception as e:
            ret = {}
            ret['ok'] = False
            ret['error'] = str(e)
        return ret

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



    def get_lista_reporte(self, carnet):
        ret = {}
        try:
            with self.connection.cursor() as cursor:
                sql = ''
                if carnet is None:
                    sql = 'SELECT r.ID_reporte, e.carnet, e.nombre, e.curso, r.procesado, r.fecha, r.mensaje FROM Estudiante e INNER JOIN Reporte r ON r.estudiante = e.ID_estudiante'
                    cursor.execute(sql)
                else:
                    sql = 'SELECT r.ID_reporte, e.carnet, e.nombre, e.curso, r.procesado, r.fecha, r.mensaje FROM Estudiante e INNER JOIN Reporte r ON r.estudiante = e.ID_estudiante WHERE e.carnet = %s'
                    cursor.execute(sql, (carnet,))

                ret = cursor.fetchall()
        except Exception as e:
            ret = {}
            ret['ok'] = False
            ret['error'] = str(e)
        return ret