# FlyerBird
Webservice for Origami instruct

## Execute
To run script, execute:

```
go run main.go

```

and open browser on http://localhost:3000

## Test (TBD)

To run specs, execute:

```
ginkgo specs/vendors

```
### run single spec
Add a `F` in front of the line to focus on. See http://onsi.github.io/ginkgo/#focused-specs

## Deploy

Before deployment, make sure to update dependencies

`godep save`

more see: http://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html


## TODO

- [x] Simple webservice for static file & pdf generation
- [x] fake pdf response
- [x] add croppper
- [ ] On crop call backendservice
- [ ] Drag&Drop File
- [ ] Duplicate workflow for Backside
- [ ] Pass query params to CLI
- [ ] integrate actual CLI (python!!)
- [ ] get python working on Heroku
- [ ] Design
- [ ] test mobile / responsive
- [ ]
- [ ]


## Sources

### Uploader
- http://codegeekz.com/best-jquery-file-upload-plugins/

#### implementation

### Croper
- http://techslides.com/image-zoom-drag-and-crop-with-html5
- http://jqueryhouse.com/jquery-image-crop-and-resize-plugins
- http://stackoverflow.com/questions/20533191/dropzone-js-client-side-image-resizing/29951416#29951416
- http://jplugins.directory/jquery-cropbox/
- http://www.jqueryrain.com/demo/jquery-crop-image-plugin

#### implementation
- http://picturecut.tuyoshi.com.br/ (with upload!!!)
- https://github.com/hongkhanh/cropbox
- http://odyniec.net/projects/imgareaselect
- https://github.com/matiasgagliano/guillotine
