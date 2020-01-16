# Script to create new lint from template

USAGE="Usage: $0 <ARG1> <ARG2> <ARG3>

ARG1: Path_name
ARG2: File_name/TestName (no 'lint_' prefix)
ARG3: Struct_name"

if [ $# -eq 0 ]; then
    echo "No arguments provided..."
    echo "$USAGE"
    exit 1
fi

if [ $# -eq 1 ]; then
    echo "Not enough arguments provided..."
    echo "$USAGE"
    exit 1
fi

if [ $# -eq 2 ]; then
    echo "Not enough arguments provided..."
    echo "$USAGE"
    exit 1
fi

if [ ! -d lints/$1 ]
then
   echo "Directory $1 does not exist. Can't make new file."
   exit 1
fi


if [ -e lints/$1/lint_$2.go ]
then
   echo "File already exists. Can't make new file."
   exit 1
fi

PATHNAME=$1
FILENAME=$2
TESTNAME=$3

sed -e "s/PACKAGE/${PATHNAME}/" \
    -e "s/SUBST/${TESTNAME}/g" \
    -e "s/SUBTEST/${FILENAME}/g" template > lints/${PATHNAME}/lint_${FILENAME}.go

echo "Created file lint_${FILENAME}.go with test name ${TESTNAME}"
