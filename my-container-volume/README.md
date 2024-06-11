
docker run -d --name my-postgres-ramadhanveha -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -v my-pg-volume-ramadhanveha:/var/lib/postgresql/data -p 5431:5432 postgres

volume :

![alt text](image.png)

container :

![alt text](image-1.png)

koneksi ke server baru :
![alt text](image-2.png)

membuat table baru :
![alt text](image-3.png)

membuat yang baru :
docker run -d --name my-postgres-v2-ramadhanveha -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -v my-pg-volume-ramadhanveha:/var/lib/postgresql/data -p 5431:5432 postgres

volume masih sama :
![alt text](image-4.png)

container :
![alt text](image-5.png)

table masih ada :
![alt text](image-6.png)
