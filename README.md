## [Install go | Arch Linux](https://go.dev/doc/install)   
> Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:   
> (You may need to run the command as root or through sudo).   
> Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.   

```
cd
cd Downloads
su root
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.3.linux-amd64.tar.gz
```

Add /usr/local/go/bin to the PATH environment variable.   
In a new terminal.   

```
cd
sudo nano ~/.xprofile
export GOROOT=/usr/local/go                                                                                        
export GOPATH=$HOME/Workspace/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

Restart.   

Install redis (optional), protobuf, protoc-gen-go, etcd, goctl, Go for VS Code (Extension), Ctrl+Shift+P (Go>Install/Update Tools).   

```
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

## [Install MariaDB | Arch Linux](https://wiki.archlinux.org/title/MySQL)   
https://wiki.archlinux.org/title/MariaDB   
> Install [MariaDB](https://archlinux.org/packages/extra/x86_64/mariadb) and additional packages (mariadb-clients, mariadb-libs) in Arch Linux.   
> Install mariadb, and run the following command before starting the mariadb.service:   

```
mariadb-install-db --user=mysql --basedir=/usr --datadir=/var/lib/mysql
```

> Now mariadb.service can be [started and/or enabled](https://mariadb.com/kb/en/systemd).   
> Starting the MariaDB Server Process on Boot.   
> MariaDB's systemd service can be configured to start at boot by executing the following:   

```
sudo systemctl enable mariadb.service
```

> Starting the MariaDB Server Process.   
> MariaDB's systemd service can be started by executing the following:   

```
sudo systemctl start mariadb.service
```

> Once you have started the MySQL server and added a root account, you may want to change the default configuration.   
> To log in as root on the MySQL server, use the following command:   

```
sudo mysql -u root -p
```

> Add user.   
> Creating a new user takes two steps: create the user; grant privileges. In the below example, the user monty with some_pass as password is being created, then granted full permissions to the database mydb:   
The reason why we do this is that the prisma tool we are going to use creates shadow database for development, and requires [shadow database user permissions](https://www.prisma.io/docs/concepts/components/prisma-migrate/shadow-database).   
> In order to create and delete the shadow database when using development commands such as migrate dev and migrate reset, Prisma Migrate currently requires that the database user defined in your datasource has permission to create databases.   
> MySQL: Database user must have CREATE, ALTER, DROP, REFERENCES ON *.* privileges.   

```
MariaDB> CREATE USER 'monty'@'localhost' IDENTIFIED BY 'some_pass';
MariaDB> GRANT ALL PRIVILEGES ON *.* TO 'monty'@'localhost';
MariaDB> FLUSH PRIVILEGES;
MariaDB> quit
```

## Set up [Prisma Client Go](https://github.com/prisma/prisma-client-go)   
https://www.prisma.io/docs/reference/database-reference/supported-databases   
https://www.prisma.io/docs/getting-started/quickstart   
> Initialise a new Go project.   
> If you don't have a Go project yet, initialise one using Go modules:   

> Init multi-module workspaces.   
https://go.dev/doc/tutorial/workspaces   

```
cd
cd Workspace
go work init ./go-demo
mkdir go-demo
cd go-demo
go mod init go-demo
```

> Get Prisma Client Go.   

```
go get -u github.com/prisma/prisma-client-go
```

> Prepare your database schema in a schema.prisma file. For example, a simple schema with a sqlite database and Prisma Client Go as a generator with two models would look like this:   

```
datasource db {
    // could be postgresql or mysql
    provider = "mysql"
    url      = "mysql://monty:password@localhost:3306/mydb"
}

generator db {
    provider = "go run github.com/prisma/prisma-client-go"
    // set the output folder and package name
    // output           = "./your-folder"
    // package          = "yourpackagename"
}

model Post {
    id        String    @id @default(cuid())
    createdAt DateTime  @default(now())
    updatedAt DateTime  @updatedAt
    title     String
    published Boolean
    desc      String?
    Comment   Comment[]
}

model Comment {
    id        String   @id @default(cuid())
    createdAt DateTime @default(now())
    content   String

    post   Post   @relation(fields: [postID], references: [id])
    postID String
}
```

> To get this up and running in your database, we use the Prisma migration tool migrate to create and migrate our database:   

```
go run github.com/prisma/prisma-client-go migrate dev --name init
go mod tidy
```

> After the migration, the Prisma Client Go client is automatically generated in your project.   

> If you just want to re-generate the client, run:   

```
go run github.com/prisma/prisma-client-go generate.
```

## Use go-zero to init an API-Gateway (Hello World Demo)   

Install go-zero, protoc-gen-go.   

```
go get -u github.com/zeromicro/go-zero@latest
go get -u google.golang.org/protobuf@latest
```

Create greet service (API).   

```
goctl api new greet --style=goZero
go mod tidy
```

Write logic:   

```
nano greet/internal/logic/greetlogic.go
```

Write the response message.   

```
diff
func (l *GreetLogic) Greet(req types.Request) (*types.Response, error) {
+   return &types.Response{
+       Message: "Hello go-zero",
+   }, nil
-   return
}
```

Start and access the service.   
Start service.   

```
cd greet
go run greet.go -f etc/greet-api.yaml
```

Access service.   

```
curl -i -X GET http://localhost:8888/from/you
```

## [Generate Protobuf from mysql and use generated prodoc to write proto file for rpc](https://github.com/Mikaelemmmm/sql2pb)   

Go install `Mikaelemmmm/sql2pb`.   

```
cd ../..
go install github.com/Mikaelemmmm/sql2pb@latest
```

Generate Protobuf from mysql.   

```
cd go-demo
mkdir service
cd service
sql2pb -go_package ./pb -host localhost -package pb -password password -port 3306 -schema mydb -service_name service -user yuzuki > service.proto
```

Directory Structure.   

```
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

Edit`post.proto`.   

```
syntax = "proto3";

option go_package ="./post";

package post;

// ------------------------------------ 
// Messages
// ------------------------------------ 


//--------------------------------Post--------------------------------
message Post {
  string id = 1; //id
  int64 createdAt = 2; //createdAt
  int64 updatedAt = 3; //updatedAt
  string title = 4; //title
  int64 published = 5; //published
  string desc = 6; //desc
}

message AddPostReq {
  int64 createdAt = 1; //createdAt
  int64 updatedAt = 2; //updatedAt
  string title = 3; //title
  int64 published = 4; //published
  string desc = 5; //desc
}

message AddPostResp {
}

message UpdatePostReq {
  string id = 1; //id
  int64 createdAt = 2; //createdAt
  int64 updatedAt = 3; //updatedAt
  string title = 4; //title
  int64 published = 5; //published
  string desc = 6; //desc
}

message UpdatePostResp {
}

message DelPostReq {
  int64 id = 1; //id
}

message DelPostResp {
}

message GetPostByIdReq {
  int64 id = 1; //id
}

message GetPostByIdResp {
  Post post = 1; //post
}

message SearchPostReq {
  int64 page = 1; //page
  int64 pageSize = 2; //pageSize
  string id = 3; //id
  int64 createdAt = 4; //createdAt
  int64 updatedAt = 5; //updatedAt
  string title = 6; //title
  int64 published = 7; //published
  string desc = 8; //desc
}

message SearchPostResp {
  repeated Post post = 1; //post
}



// ------------------------------------ 
// Rpc Func
// ------------------------------------ 

service service{ 

	 //-----------------------Post----------------------- 
	 rpc AddPost(AddPostReq) returns (AddPostResp); 
	 rpc UpdatePost(UpdatePostReq) returns (UpdatePostResp); 
	 rpc DelPost(DelPostReq) returns (DelPostResp); 
	 rpc GetPostById(GetPostByIdReq) returns (GetPostByIdResp); 
	 rpc SearchPost(SearchPostReq) returns (SearchPostResp); 

}
```

Edit`comment.proto`.   

```
syntax = "proto3";

option go_package ="./comment";

package comment;

// ------------------------------------ 
// Messages
// ------------------------------------ 

//--------------------------------Comment--------------------------------
message Comment {
  string id = 1; //id
  int64 createdAt = 2; //createdAt
  string content = 3; //content
  string postID = 4; //postID
}

message AddCommentReq {
  int64 createdAt = 1; //createdAt
  string content = 2; //content
  string postID = 3; //postID
}

message AddCommentResp {
}

message UpdateCommentReq {
  string id = 1; //id
  int64 createdAt = 2; //createdAt
  string content = 3; //content
  string postID = 4; //postID
}

message UpdateCommentResp {
}

message DelCommentReq {
  int64 id = 1; //id
}

message DelCommentResp {
}

message GetCommentByIdReq {
  int64 id = 1; //id
}

message GetCommentByIdResp {
  Comment comment = 1; //comment
}

message SearchCommentReq {
  int64 page = 1; //page
  int64 pageSize = 2; //pageSize
  string id = 3; //id
  int64 createdAt = 4; //createdAt
  string content = 5; //content
  string postID = 6; //postID
}

message SearchCommentResp {
  repeated Comment comment = 1; //comment
}



// ------------------------------------ 
// Rpc Func
// ------------------------------------ 

service service{ 

	 //-----------------------Comment----------------------- 
	 rpc AddComment(AddCommentReq) returns (AddCommentResp); 
	 rpc UpdateComment(UpdateCommentReq) returns (UpdateCommentResp); 
	 rpc DelComment(DelCommentReq) returns (DelCommentResp); 
	 rpc GetCommentById(GetCommentByIdReq) returns (GetCommentByIdResp); 
	 rpc SearchComment(SearchCommentReq) returns (SearchCommentResp); 

}
```

## [Generate Model from mysql, Generate Rpc using Prodoc](https://github.com/Mikaelemmmm/sql2pb)   

```
goctl model mysql datasource -url="yuzuki:password@tcp(localhost:3306)/mydb" -table="Post" -dir=./post/model --style=goZero
goctl model mysql datasource -url="yuzuki:password@tcp(localhost:3306)/mydb" -table="Comment" -dir=./comment/model --style=goZero
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

## Business Coding (Coding rpc)   

Add database configuration for rpc in yaml file.   

```
nano ./post/rpc/etc/post.yaml
```

Edit`post.yaml`.   

```diff
  Key: post.rpc
+ Mysql:
+   DataSource: $user:$password@tcp($url)/$db?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo
```

Add database struct for rpc config.   

```
nano ./post/rpc/internal/config/config.go
```

Edit`config.go`.   

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

```
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

Add resource dependency (srvice context or model) for rpc in svc:  

```
nano ./post/rpc/internal/svc/serviceContext.go
```

Edit`serviceContext.go`.   

```diff
package svc

import (
	"go-demo/service/post/model"
+	"go-demo/service/post/rpc/internal/config"

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

```
nano ./post/rpc/internal/logic/getPostByIdLogic.go
```

Edit`getPostByIdLogic.go`.   

```diff
func (l *GetPostByIdLogic) GetPostById(in *post.GetPostByIdReq) (*post.GetPostByIdResp, error) {
	// todo: add your logic here and delete this line
-   return &post.GetPostByIdResp{}, nil
+	return &post.GetPostByIdResp{
+		Post: &post.Post{
+			Id:    "1",
+			Title: "TestPost",
+		},
+	}, nil
}
```

## Coding APIs   

Define APIs:   

```
touch ./post/api/postApi.api
```

Edit`postApi.api`.   

```
type (
	GetPostByIdReq {
		Id int64 `json:"id"`
	}

	GetPostByIdResp {
		Post *Post `json:"post"`
	}

	Post {
		Id        string `json:"id"`
		CreatedAt int64  `json:"createdAt"`
		UpdatedAt int64  `json:"updatedAt"`
		Title     string `json:"title"`
		Published int64  `json:"published"`
		Desc      string `json:"desc"`
	}
)

service user-api {
	@handler post
	get /api/post/get/:id (GetPostByIdReq) returns (GetPostByIdResp)
}
```

Generate api services.   

```
goctl api go --api ./post/api/postApi.api -dir ./post/api --style=goZero
```

Add restapi configuration for api config.   

```
nano ./post/api/internal/config/config.go
```

Edit`config.go`.   

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
+	PostRpc zrpc.RpcServerConf
}
```

Add yaml configuration.   

```
nano ./post/api/etc/post-api.yaml
```

Intergrate rpc for api in yaml file.   

```diff
Port: 8888
+ PostRpc:
+   Etcd:
+     Hosts:
+     - 127.0.0.1:2379
+     Key: post.rpc
```

Refine the service dependencies:

```
nano ./post/internal/svc/
```

To test the post-api:

```
etcd

```