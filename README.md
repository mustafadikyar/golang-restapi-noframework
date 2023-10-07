# golang-restapi-noframework-postgres

Bu proje, **Go (Golang)** dilini kullanarak basit bir **REST API** uygulamasını içermektedir. 

**PostgreSQL** veritabanını kullanarak ürünlerle ilgili temel *CRUD (Create, Read, Update, Delete)* operasyonları gerçekleştirmek mümkündür. 

Projede bir web framework kullanılmadan, sadece Go'nun standart kütüphaneleri ve PostgreSQL veritabanı sorgularıyla veritabanı işlemleri yönetilmiştir.

## API Endpointleri

### 1. Ürün Oluşturmak

**Endpoint:** `POST /products`

Bu endpoint, yeni bir ürün oluşturmak için kullanılır. JSON formatında bir ürün nesnesi alır ve veritabanına ekler.

```json
Örnek İstek:
{
  "name": "Yeni Ürün"
}
```



### 2. Tüm Ürünleri Listeleme

**Endpoint:** `GET /products`

Bu endpoint, tüm ürünleri listeleyen bir GET isteği için kullanılır. Veritabanındaki tüm ürünleri döndürür.

### 3. Belirli Bir Ürünü Görüntüleme

**Endpoint:** `GET /products/{id}`

Bu endpoint, belirli bir ürünün detaylarını görmek için kullanılır. {id} yerine ürünün benzersiz kimliği (ID) konmalıdır.

### 4. Ürün Güncelleme

**Endpoint:** `PUT /products/{id}`

Bu endpoint, belirli bir ürünün bilgilerini güncellemek için kullanılır. {id} yerine güncellenmek istenen ürünün ID'si konmalıdır.

```json
Örnek İstek:
{
  "name": "Güncellenmiş Ürün Adı"
}
```

### 5. Ürün Silme

**Endpoint:** `DELETE /products/{id}`

Bu endpoint, belirli bir ürünü silmek için kullanılır. {id} yerine silinmek istenen ürünün ID'si konmalıdır.
