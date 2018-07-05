#! /bin/bash 
if [[ ${1} == "--clean" ]] ; then
    clean=1
fi
continueOnError=
# A small utility to build fixture servers
# Fixtures with models only
testcases="${testcases} ./fixture-1601.yaml"

# Generation options
fullFlatten="--with-flatten=full"
withExpand="--with-expand"
minimal="--with-flatten=minimal"
for opts in ${fullFlatten} ${withExpand} ${minimal} ; do
for testcase in ${testcases} ; do
    grep -q discriminator ${testcase}
    discriminated=$?
    if [[ ${discriminated} -eq 0 && ${opts} == ${withExpand} ]] ; then
        echo "Skipped ${testcase} with ${opts}: discriminator not supported with ${opts}"
        continue
    fi
    if [[ ${testcase} == "../1479/fixture-1479-part.yaml" && ${opts} == ${withExpand} ]] ; then
        echo "Skipped ${testcase} with ${opts}: known issue with enum in anonymous allOf not validated. See you next PR"
        continue
    fi

    spec=${testcase}
    testcase=`basename ${testcase}`
    if [[ -z ${opts} || ${opts} == ${minimal} ]] ; then
        target=./gen-${testcase%.*}-minimal
    elif [[ ${opts} ==  ${fullFlatten} ]] ; then
        target=./gen-${testcase%.*}-flatten
    else
        target=./gen-${testcase%.*}-expand
    fi
    serverName="codegensrv"
    rm -rf ${target}
    mkdir ${target}
    echo "Model generation for ${spec} with opts=${opts}"
    swagger generate model --skip-validation ${opts} --spec=${spec} --target=${target} --output=${testcase%.*}.log
    # 1>x.log 2>&1
    #
    if [[ $? != 0 ]] ; then
        echo "Generation failed for ${spec}"
        if [[ ! -z ${continueOnError} ]] ; then
            failures="${failures} codegen:${spec}"
            continue
        else
            exit 1
        fi
    fi
    echo "${spec}: Generation OK"
    if [[ ! -d ${target}/models ]] ; then
        echo "No model in this spec! Skipped"
    else
        (cd ${target}/models; go build)
        if [[ $? != 0 ]] ; then
            echo "Build failed for ${spec}"
            if [[ ! -z ${continueOnError} ]] ; then
                failures="${failures} build:${spec}"
                continue
            else
                exit 1
            fi
        fi
        echo "${spec}: Build OK"
        if [[ -n ${clean} ]] ; then 
             rm -rf ${target}
        fi
    fi
done
done
if [[ ! -z ${failures} ]] ; then 
    echo ${failures}|tr ' ' '\n'
else
    echo "No failures"
fi
exit
