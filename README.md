<div id="top"></div>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/yuzukicat/go-zero-demo">
    <img src="./assets/getPostById.png" alt="go-zero codegen">
  </a>
  </br>
  </br>
  <a href="https://github.com/yuzukicat/go-zero-demo">
    <img src="./assets/PostTable.png" alt="go-zero codegen">
  </a>
  </br>
  </br>
  <h3 align="center">Use go-zero, prisma to introspect SQL, generating go struct, protobuf, gRPC server / client and REST API</h3>
  </br>
  <p align="left">
    This article is going to build a working demo using go-zero framework, which is a golang api framework using grpc and restapi. It supports codegeration. Prisma is also used to prototyping a database and seed the database. An SQL gen go tool is used to codegen protobuf files, which will be used as templates to generate rpc and api with full-stack frameworks and auto generating graphQL crud operations. The system is arch linux and the database is MariaDB.
    <br />
    <br />
    <a href="https://github.com/yuzukicat/go-zero-demo"><strong>The source code for the working demo in github »</strong></a>
    .
    <a href="https://github.com/yuzukicat/parkinsons-disease/issues"><strong>The blog article  »</strong></a>
  </p>
</div>

## Install go | Arch Linux [1] [2]  

> Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:   

> (You may need to run the command as root or through sudo).   

> Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.   

```sh
cd
cd Downloads
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.3.linux-amd64.tar.gz
```

In a new terminal.   

```sh
cd
sudo nano ~/.xprofile
export GOROOT=/usr/local/go                                                                                        
export GOPATH=$HOME/Workspace/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

Restart.   

Install ```redis (optional), protobuf, protoc-gen-go, etcd, goctl, Go for VS Code (Extension), Ctrl+Shift+P (Go>Install/Update Tools), ext:proto```.   

```sh
cd Package
git clone https://aur.archlinux.org/protoc-gen-go.git
cd protoc-gen-go
makepkg -i
cd ..
git clone https://aur.archlinux.org/etcd.git
cd etcd
makepkg -i
cd
cd Workspace
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

<!-- ABOUT THE PROJECT -->
## Install MariaDB | Arch Linux [3] [4]   

> Install ```MariaDB``` [5] and additional packages (```mariadb-clients, mariadb-libs```) in Arch Linux.   

> Install mariadb, and run the following command ```before starting the mariadb.service```:   

```sh
mariadb-install-db --user=mysql --basedir=/usr --datadir=/var/lib/mysql
```

> Now mariadb.service can be ```started and/or enabled``` [6].   

> Starting the MariaDB Server Process on Boot.   

> MariaDB's systemd service can be configured to start at boot by executing the following:   

```sh
sudo systemctl enable mariadb.service
```

> Starting the MariaDB Server Process.   

> MariaDB's systemd service can be started by executing the following:   

```sh
sudo systemctl start mariadb.service
```

> Once you have started the MySQL server and added a root account, you may want to change the default configuration.   

> To log in as root on the MySQL server, use the following command:   

```sh
sudo mysql -u root -p
```

> Add user.   

> Creating a new user takes two steps: create the user; grant privileges. In the below example, the user monty with some_pass as password is being created, then ```granted full permissions``` to the database mydb:   
The reason why we do this is that the prisma tool we are going to use creates shadow database for development, and requires ```shadow database user permissions``` [7].   

> In order to create and delete the shadow database when using development commands such as migrate dev and migrate reset, Prisma Migrate currently requires that the database user defined in your datasource has permission to create databases.   

> MySQL: Database user must have CREATE, ALTER, DROP, REFERENCES ON *.* privileges.   

```sql
MariaDB> CREATE USER 'monty'@'localhost' IDENTIFIED BY 'some_pass';
MariaDB> GRANT ALL PRIVILEGES ON *.* TO 'monty'@'localhost';
MariaDB> FLUSH PRIVILEGES;MariaDB
```

> Initialise a new Go project.   

> If you don't have a Go project yet, initialise one using Go modules:   

> Init multi-module workspaces [8].   

```sh
cd
cd Workspace
go work init ./go-demo
mkdir go-demo
cd go-demo
go mod init go-demo
```

> Get Prisma Client Go.   

```sh
go get -u github.com/prisma/prisma-client-go
```

> Prepare your database schema in a schema.prisma file. For example, a simple schema with a sqlite database and Prisma Client Go as a generator with two models would look like this:   

```
datasource db {
    // could be postgresql or mysql
    provider = "mysql"
    url      = "mysql://username:password@localhost:3306/mydb"
}

generator db {
    provider = "go run github.com/prisma/prisma-client-go"
    // set the output folder and package name
    // output           = "./your-folder"
    // package          = "yourpackagename"
}

model Post {
    id        Int       @id @default(autoincrement())
    createdAt DateTime  @default(now())
    updatedAt DateTime  @updatedAt
    title     String
    published Boolean
    desc      String?
    Comment   Comment[]
}

model Comment {
    id        Int      @id @default(autoincrement())
    createdAt DateTime @default(now())
    content   String

    post   Post @relation(fields: [postID], references: [id])
    postID Int
}
```

> To get this up and running in your database, we use the Prisma migration tool migrate to create and migrate our database:   

```sh
go run github.com/prisma/prisma-client-go migrate dev --name init
go mod tidy
```

> After the migration, the Prisma Client Go client is automatically generated in your project.   

Seed mydb>Post table with test data.

Edit `migrations/xxxx_init/migration.sql`.   

```diff
+ -- SEED
+ INSERT INTO `Post` (`id`,`updatedAt`,`title`,`published`, `desc`) VALUES (1,CURRENT_TIMESTAMP(3),"testpost",false, "desc");
```

> If you just want to re-generate the client, run:   

```sh
go run github.com/prisma/prisma-client-go generate.
```

<!-- ABOUT THE PROJECT -->
## Use go-zero to init an API-Gateway (Hello World Demo) [9]   

Install go-zero, protoc-gen-go.   

```sh
go get -u github.com/zeromicro/go-zero@latest
go get -u google.golang.org/protobuf@latest
```

Create greet service (API).   

```sh
goctl api new greet --style=goZero
go mod tidy
```

Write logic:   

```sh
nano greet/internal/logic/greetlogic.go
```

Write the response message.   

```diff
func (l *GreetLogic) Greet(req types.Request) (*types.Response, error) {
+   return &types.Response{
+       Message: "Hello go-zero",
+   }, nil
-   return
}
```

Start and access the service.   
Start service.   

```sh
cd greet
go run greet.go -f etc/greet-api.yaml
```

Access service.   

```sh
curl -i -X GET http://localhost:8888/from/you
```

<!-- ABOUT THE PROJECT -->
## Generate Protobuf from mysql and use generated prodoc to write proto file for rpc [10]   

Go install `Mikaelemmmm/sql2pb` [11].   

```sh
cd ../..
go install github.com/Mikaelemmmm/sql2pb@latest
```

Generate Protobuf from mysql.   

```sh
cd go-demo
mkdir service
cd service
sql2pb -go_package ./pb -host localhost -package pb -password password -port 3306 -schema mydb -service_name service -user username > service.proto
```

Directory Structure.   

```sh
mkdir post
mkdir comment
cd post
mkdir api
mkdir rpc
cd rpc
touch post.proto
cd ../..
cd comment
mkdir api
mkdir rpc
cd rpc
touch comment.proto
cd ../..
```

```
.
├── db
│   ├── db_gen.go
│   ├── query-engine-debian-openssl-3.0.x_gen.go
│   └── query-engine-linux_gen.go
├── go.mod
├── go.sum
├── greet
│   ├── etc
│   │   └── greet-api.yaml
│   ├── greet.api
│   ├── greet.go
│   └── internal
│       ├── config
│       │   └── config.go
│       ├── handler
│       │   ├── greetHandler.go
│       │   └── routes.go
│       ├── logic
│       │   └── greetLogic.go
│       ├── svc
│       │   └── serviceContext.go
│       └── types
│           └── types.go
├── migrations
│   ├── 20221116081516_init
│   │   └── migration.sql
│   └── migration_lock.toml
├── README.md
├── schema.prisma
└── service
    ├── comment
    │   ├── api
    │   ├── model
    │   │   ├── commentModel_gen.go
    │   │   ├── commentModel.go
    │   │   └── vars.go
    │   └── rpc
    │       ├── comment.go
    │       ├── comment.proto
    │       ├── etc
    │       │   └── comment.yaml
    │       ├── internal
    │       │   ├── config
    │       │   │   └── config.go
    │       │   ├── logic
    │       │   │   ├── addCommentLogic.go
    │       │   │   ├── delCommentLogic.go
    │       │   │   ├── getCommentByIdLogic.go
    │       │   │   ├── searchCommentLogic.go
    │       │   │   └── updateCommentLogic.go
    │       │   ├── server
    │       │   │   └── serviceServer.go
    │       │   └── svc
    │       │       └── serviceContext.go
    │       ├── service
    │       │   └── service.go
    │       └── types
    │           └── comment
    │               ├── comment_grpc.pb.go
    │               └── comment.pb.go
    ├── default.etcd
    │   └── member
    │       ├── snap
    │       │   └── db
    │       └── wal
    │           └── 0000000000000000-0000000000000000.wal
    ├── post
    │   ├── api
    │   │   ├── etc
    │   │   │   └── post-api.yaml
    │   │   ├── internal
    │   │   │   ├── config
    │   │   │   │   └── config.go
    │   │   │   ├── handler
    │   │   │   │   ├── getPostByIdHandler.go
    │   │   │   │   └── routes.go
    │   │   │   ├── logic
    │   │   │   │   └── getPostByIdLogic.go
    │   │   │   ├── svc
    │   │   │   │   └── serviceContext.go
    │   │   │   └── types
    │   │   │       └── types.go
    │   │   ├── postApi.api
    │   │   └── post.go
    │   ├── model
    │   │   ├── postModel_gen.go
    │   │   ├── postModel.go
    │   │   └── vars.go
    │   └── rpc
    │       ├── etc
    │       │   └── post.yaml
    │       ├── internal
    │       │   ├── config
    │       │   │   └── config.go
    │       │   ├── logic
    │       │   │   ├── addPostLogic.go
    │       │   │   ├── delPostLogic.go
    │       │   │   ├── getPostByIdLogic.go
    │       │   │   ├── searchPostLogic.go
    │       │   │   └── updatePostLogic.go
    │       │   ├── server
    │       │   │   └── serviceServer.go
    │       │   └── svc
    │       │       └── serviceContext.go
    │       ├── post.go
    │       ├── post.proto
    │       ├── service
    │       │   └── service.go
    │       └── types
    │           └── post
    │               ├── post_grpc.pb.go
    │               └── post.pb.go
    └── service.proto
```

Edit `post.proto`.   

```
syntax = "proto3";

option go_package ="./post";

package post;

// ------------------------------------ 
// Messages
// ------------------------------------ 


//--------------------------------Post--------------------------------
syntax = "proto3";

option go_package = "./post";

package post;

// ------------------------------------
// Messages
// ------------------------------------

//--------------------------------Post--------------------------------
message Post {
  int64 id = 1;        // id
  int64 createdAt = 2; // createdAt
  int64 updatedAt = 3; // updatedAt
  string title = 4;    // title
  int64 published = 5; // published
  string desc = 6;     // desc
}

message AddPostReq {
  int64 createdAt = 1; // createdAt
  int64 updatedAt = 2; // updatedAt
  string title = 3;    // title
  int64 published = 4; // published
  string desc = 5;     // desc
}

message AddPostResp {}

message UpdatePostReq {
  int64 id = 1;        // id
  int64 createdAt = 2; // createdAt
  int64 updatedAt = 3; // updatedAt
  string title = 4;    // title
  int64 published = 5; // published
  string desc = 6;     // desc
}

message UpdatePostResp {}

message DelPostReq {
  int64 id = 1; // id
}

message DelPostResp {}

message GetPostByIdReq {
  int64 id = 1; // id
}

message GetPostByIdResp {
  Post post = 1; // post
}

message SearchPostReq {
  int64 page = 1;      // page
  int64 pageSize = 2;  // pageSize
  int64 id = 3;        // id
  int64 createdAt = 4; // createdAt
  int64 updatedAt = 5; // updatedAt
  string title = 6;    // title
  int64 published = 7; // published
  string desc = 8;     // desc
}

message SearchPostResp {
  repeated Post post = 1; // post
}

// ------------------------------------
// Rpc Func
// ------------------------------------

service service {

  //-----------------------Post-----------------------
  rpc AddPost(AddPostReq) returns (AddPostResp);
  rpc UpdatePost(UpdatePostReq) returns (UpdatePostResp);
  rpc DelPost(DelPostReq) returns (DelPostResp);
  rpc GetPostById(GetPostByIdReq) returns (GetPostByIdResp);
  rpc SearchPost(SearchPostReq) returns (SearchPostResp);
}
```

Edit `comment.proto`.   

```
syntax = "proto3";

option go_package = "./comment";

package comment;

// ------------------------------------
// Messages
// ------------------------------------

//--------------------------------Comment--------------------------------
message Comment {
  int64 id = 1;        // id
  int64 createdAt = 2; // createdAt
  string content = 3;  // content
  int64 postID = 4;    // postID
}

message AddCommentReq {
  int64 createdAt = 1; // createdAt
  string content = 2;  // content
  int64 postID = 3;    // postID
}

message AddCommentResp {}

message UpdateCommentReq {
  int64 id = 1;        // id
  int64 createdAt = 2; // createdAt
  string content = 3;  // content
  int64 postID = 4;    // postID
}

message UpdateCommentResp {}

message DelCommentReq {
  int64 id = 1; // id
}

message DelCommentResp {}

message GetCommentByIdReq {
  int64 id = 1; // id
}

message GetCommentByIdResp {
  Comment comment = 1; // comment
}

message SearchCommentReq {
  int64 page = 1;      // page
  int64 pageSize = 2;  // pageSize
  int64 id = 3;        // id
  int64 createdAt = 4; // createdAt
  string content = 5;  // content
  int64 postID = 6;    // postID
}

message SearchCommentResp {
  repeated Comment comment = 1; // comment
}

// ------------------------------------
// Rpc Func
// ------------------------------------

service service {

  //-----------------------Comment-----------------------
  rpc AddComment(AddCommentReq) returns (AddCommentResp);
  rpc UpdateComment(UpdateCommentReq) returns (UpdateCommentResp);
  rpc DelComment(DelCommentReq) returns (DelCommentResp);
  rpc GetCommentById(GetCommentByIdReq) returns (GetCommentByIdResp);
  rpc SearchComment(SearchCommentReq) returns (SearchCommentResp);
}
```

<!-- ABOUT THE PROJECT -->
## Generate Model from mysql, Generate Rpc using Prodoc   

```sh
goctl model mysql datasource -url="username:password@tcp(localhost:3306)/mydb" -table="Post" -dir=./post/model --style=goZero
goctl model mysql datasource -url="username:password@tcp(localhost:3306)/mydb" -table="Comment" -dir=./comment/model --style=goZero
cd ..
go mod tidy
cd service
goctl rpc protoc ./post/rpc/post.proto --go_out=./post/rpc/types --go-grpc_out=./post/rpc/types --zrpc_out=./post/rpc --style=goZero
goctl rpc protoc ./comment/rpc/comment.proto --go_out=./comment/rpc/types --go-grpc_out=./comment/rpc/types --zrpc_out=./comment/rpc --style=goZero
sed -i 's/,omitempty//g' ./post/rpc/types/post/post.pb.go
sed -i 's/,omitempty//g' ./comment/rpc/types/comment/comment.pb.go
cd ..
go mod tidy
cd service
```

<!-- ABOUT THE PROJECT -->
## Business Coding (Coding rpc)   

Add database configuration for rpc in yaml file.   

```sh
nano ./post/rpc/etc/post.yaml
```

Edit `post.yaml`.   

```diff
  Key: post.rpc
+Mysql:
+  DataSource: username:password@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo
```

Add database struct for rpc config.   

```sh
nano ./post/rpc/internal/config/config.go
```

Edit `config.go`.   

```diff
package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
+ 	Mysql struct {
+		DataSource string
+	}
}
```

In model, comment unused datasource.   

```sh
nano ./post/model/postModel.go
```

Edit`postModel.go`.   

```diff
import (
- 	"github.com/lib/pq"
+	// "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)
```

Also edit `commentModel.go` in the same way.   

Add resource dependency (srvice context or model) for rpc in svc:  

```sh
nano ./post/rpc/internal/svc/serviceContext.go
```

Edit `serviceContext.go`.   

```diff
package svc

import (
+	"go-demo/service/post/model"
	"go-demo/service/post/rpc/internal/config"

+	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
+	PostModel model.PostModel
}

func NewServiceContext(c config.Config) *ServiceContext {
+	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
+		PostModel: model.NewPostModel(conn),
	}
}
```

Write logic for rpc.   

```sh
nano ./post/rpc/internal/logic/getPostByIdLogic.go
```

Edit `getPostByIdLogic.go`.   

```diff
func (l *GetPostByIdLogic) GetPostById(in *post.GetPostByIdReq) (*post.GetPostByIdResp, error) {
	// todo: add your logic here and delete this line
+	onePost, err := l.svcCtx.PostModel.FindOne(l.ctx, in.Id)
+	switch err {
+	case nil:
+	case model.ErrNotFound:
+		return nil, errors.New("post does not exist")
+	default:
+		return nil, err
+	}
- return &post.GetPostByIdResp{}, nil
+	return &post.GetPostByIdResp{
+		Post: &post.Post{
+			Id:        onePost.Id,
+			CreatedAt: onePost.CreatedAt.Unix(),
+			UpdatedAt: onePost.UpdatedAt.Unix(),
+			Title:     onePost.Title,
+			Published: onePost.Published,
+			Desc:      onePost.Desc,
+		},
+	}, nil
}
```

<!-- ABOUT THE PROJECT -->
## Generating APIs   

Define APIs:   

```sh
touch ./post/api/postApi.api
```

Edit `postApi.api`.   

```go
type (
	GetPostByIdReq {
		Id int64 `json:"id"`
	}

	GetPostByIdResp {
		Post *Post `json:"post"`
	}

	Post {
		Id        int64  `json:"id"`
		CreatedAt int64  `json:"createdAt"`
		UpdatedAt int64  `json:"updatedAt"`
		Title     string `json:"title"`
		Published int64  `json:"published"`
		Desc      string `json:"desc"`
	}
)

service post-api {
	@handler getPostById
	post /api/post/getPostById (GetPostByIdReq) returns (GetPostByIdResp)
}
```

Generate api services.   

```sh
goctl api go --api ./post/api/postApi.api -dir ./post/api --style=goZero
```

Edit `config.go`.   

```diff
package config

- import "github.com/zeromicro/go-zero/zrpc"
+ import (
+	"github.com/zeromicro/go-zero/rest"
+	"github.com/zeromicro/go-zero/zrpc"
+)

type Config struct {
+	rest.RestConf
-   zrpc.RpcServerConf
+	PostRpc zrpc.RpcClientConf
}
```

Add yaml configuration.   

```sh
nano ./post/api/etc/post-api.yaml
```



Intergrate rpc for api in yaml file.   

```diff
Port: 8888
+ PostRpc:
+  Etcd:
+    Hosts:
+      - 127.0.0.1:2379
+    Key: post.rpc
```

Refine the service dependencies.   

```sh
nano ./post/api/internal/svc/serviceContext.go
```

Edit`serviceContext.go`.   

```diff
package svc

import (
	"go-demo/service/post/api/internal/config"
+	"go-demo/service/post/rpc/service"

+	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
+	PostRpc service.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,https://wiki.archlinux.org/title/MySQL
import (
	"context"
+	"encoding/json"
+	"fmt"

	"go-demo/service/post/api/internal/svc"
	"go-demo/service/post/api/internal/types"
	"go-demo/service/post/rpc/types/post"

	"github.com/zeromicro/go-zero/core/logx"
)

...

func (l *PostLogic) Post(req *types.GetPostByIdReq) (resp *types.GetPostByIdResp, err error) {
	// todo: add your logic here and delete this line
- return
+	// use post rpc
+	onePost, err := l.svcCtx.PostRpc.GetPostById(l.ctx, &post.GetPostByIdReq{
+		Id: req.Id,
+	})
+	if err != nil {
+		return nil, err
+	}
+
+	return &types.GetPostByIdResp{
+		Post: &types.Post{
+			Id:        onePost.Post.Id,(https://wiki.archlinux.org/title/MySQL)
```

<!-- ABOUT THE PROJECT -->
## Known Issues

- An mySQL record can not contain ```null``` value, if null, there will be an error that can not convert null to string, which means in prisma schema, the string? should be string.
- In prisma schema, the Id should be the type Int, not string, which means that the advanced feature like cuid support is missing. Otherwise, in the generated model, the type of Id would be string (because go-zero introspacted the database and infered the type according to the database, but the protobuf uses `Mikaelemmmm/sql2pb` codegen)

<!-- ABOUT THE PROJECT -->
## References   

[1]: https://go.dev/doc/install   

[2]: https://go.dev/doc/tutorial/workspaces   

[3]: https://wiki.archlinux.org/title/MySQL   

[4]: https://wiki.archlinux.org/title/MariaDB   

[5]: https://archlinux.org/packages/extra/x86_64/mariadb   

[6]: https://mariadb.com/kb/en/systemd   

[7]: https://www.prisma.io/docs/concepts/components/prisma-migrate/shadow-database   

[8]: https://go.dev/doc/tutorial/workspaces   

[9]: https://go-zero.dev   

[10]: https://blog.csdn.net/wanglei19891210/article/details/124420212   

[11]: https://github.com/Mikaelemmmm/sql2pb   

```js
[1]: https://go.dev/doc/install
[2]: https://go.dev/doc/tutorial/workspaces
[3]: https://wiki.archlinux.org/title/MySQL
[4]: https://wiki.archlinux.org/title/MariaDB
[5]: https://archlinux.org/packages/extra/x86_64/mariadb
[6]: https://mariadb.com/kb/en/systemd
[7]: https://www.prisma.io/docs/concepts/components/prisma-migrate/shadow-database
[8]: https://go.dev/doc/tutorial/workspaces
[9]: https://go-zero.dev
[10]: https://blog.csdn.net/wanglei19891210/article/details/124420212
[11]: https://github.com/Mikaelemmmm/sql2pb
```