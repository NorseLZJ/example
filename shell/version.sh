#!/bin/env bash

s_val1='999'
s_val2='1'
s_val3='1'
s_val4='1'

i_val1=$((10#${s_val1}))
i_val2=$((10#${s_val2}))
i_val3=$((10#${s_val3}))
i_val4=$((10#${s_val4}))

echo ${i_val1}
echo ${i_val2}
echo ${i_val3}
echo ${i_val4}

if [ ${i_val1} == 999 ]; then
    echo "val1 == 999"
fi
