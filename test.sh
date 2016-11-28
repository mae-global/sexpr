echo 'build fuzzing'
go-fuzz-build pinhole/internal/sexpr
mkdir -p workdir/corpus
echo 'fuzzing...'
go-fuzz -bin=fuzz.zip -workdir=workdir


