npm=/home/arthur/soft/node-v8.12.0-linux-x64/bin/npm
config2vm1:
	scp ./engine/config/app.yml vm1:/home/arthur/app/config
engine2vm1:
	scp ./engine/engine vm1:/home/arthur/app
dist2vm1: dist2vm1copy
	ssh vm1 'cd /home/arthur/app && unzip dist_f.zip -d ./views/front && unzip dist_a.zip -d ./views/admin' 
dist2vm1copy:
	scp ./dist_a.zip ./dist_f.zip vm1:/home/arthur/app
distzip_f: distgen_f
	cd ./views/front && zip -r ../../dist_f.zip ./dist
distzip_a: distgen_a
	cd ./views/admin && zip -r ../../dist_a.zip ./production
distzip_a: distgen_a
	cd ./views/prodfront && zip -r ../../dist_pf.zip ./production
distgen_f:
	cd ./views/front && ${npm} run generate
distgen_a:
	cd ./views/admin && ${npm} run generate
lincompile:
	env GOOS=linux GOARCH=amd64 go build ./engine -o ./engine/cmd/linengine
wincompile:
	env GOOS=windows GOARCH=amd64 go build ./engine -o ./engine/cmd/winengine
