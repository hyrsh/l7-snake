#!/bin/bash


template=./tmpl_deployment.yml

name=$1

if [ "$name" == "" ]; then
  echo "No name given ..."
  postfix=$(tr -cd '[:digit:]' < /dev/urandom | fold -w4 | head -n1)
  name="snake-$postfix"
  echo "I name you $name"
fi


cat $template | sed -e "s/{{ name }}/$name/g" -e "s/{{ route }}/default/g" -e "s/{{ term }}/yes/g" -e "s/{{ target }}/127.0.0.1/g" > ${name}-dpl.yml

echo "Created deployment ${name}-dpl.yml"
