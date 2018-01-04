#!/bin/bash

i=0

while [ $i -lt 100 ]; do
    (( i++ ));
    curl http://localhost:8081&
done
