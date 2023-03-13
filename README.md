# Talks

## Tasks

### build

requires: clean

```
mkdir dist
for pathname in $( find . -maxdepth 3 -name package.json -type f -prune ); do
	p=$(echo $pathname | cut -d'/' -f2)
	echo $p
	cd $p
	yarn
	yarn run build
	cd ..
	cp -r ./$p/dist ./dist/$p
done
```

### clean

```
rm -rf dist
for pathname in $( find . -maxdepth 3 -name package.json -type f -prune ); do
	p=$(echo $pathname | cut -d'/' -f2)
	echo $p
	rm -rf ./$p/dist
done
```
