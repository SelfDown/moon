rm -rf dist/*
go build -o dist/bin -v -x
cp -r static dist
cp -r conf dist
cp -r collect dist
