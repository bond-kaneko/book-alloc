# book-alloc docs

開発用ドキュメント

# ER図

![er](book-alloc-er.png)
https://lucid.app/lucidchart/b1eb7f71-a833-4fca-9364-c45dcceb49c9/edit?viewport_loc=-427%2C22%2C2261%2C1018%2C_X9Ez5NB4sri&invitationId=inv_89088f8d-dec1-4ed4-8e75-0c96c5694980

# ディレクトリ構造

参考
https://blog.devgenius.io/golang-apis-a-skeleton-for-your-future-projects-a082dc4d6818

# localでのテストログイン

```
curl -c /tmp/cookie.txt -X POST -d '{"email": "test@example.com", "password": "password"}' localhost:8888/login
curl -b /tmp/cookie.txt localhost:8888/v1/users 
```
