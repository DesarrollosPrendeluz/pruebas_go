Ejemplo para probar conexión a bbdd con go y usar Gin como framework para apiRest

Para levantar la BBDD utiliza los siguientes comandos: 
Descargar y levantar el contenedor de mysql

```
docker run -d --name mysql-container -e TZ=UTC -p 30306:3306 -e MYSQL_ROOT_PASSWORD=1234 ubuntu/mysql:8.0-22.04_beta
```

Despues de conectarse por terminal y crear una bbdd con el nombre que queremos hacemos el dump 

```
 docker exec -i mysql-container mysql -uuser -ppassword name_db < data.sql  
```
