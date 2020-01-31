# epub-to-kindle

Convert and send epubs to a kindle device 

Just put your @kindle.com address and drop epubs 


## Build 

Get deps

```bash
go get
```

Then build

```bash
go build
```



## Deps

You will need `ebook-convert` binary from calibre. 
So install calibre on your platform. https://calibre-ebook.com/



## Using 


### CLI Mode 

Convert a sent one epub from console

```bash
./epub-to-kindle your-happy.epub my-mail@kindle.com 
```

### Web interface

Start a web server where you can drag & drop your epubs and magically those be available on your kindle device 

```bash
./epub-to-kindle -web 
```

and then open your browser and go to http://localhost:8081



### Config

You have to create your own `config.yml` based on the `config.example`
Then put your SMTP credentials and make sure the 
path to `ebook-convert` binary is right for your platform







