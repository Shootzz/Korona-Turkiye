#!/bin/bash

vaka=$(curl -s https://covid19.saglik.gov.tr/covid19api?getir=liste | cut -d "," -f3 | cut -d '"' -f4)

gun=$(curl -s https://covid19.saglik.gov.tr/covid19api?getir=liste | cut -d "." -f1 | grep -o '..$')

test_sayi=$(curl -s https://covid19.saglik.gov.tr/covid19api?getir=liste | cut -d "," -f2 | cut -d '"' -f4)

iyilesen=$(curl -s https://covid19.saglik.gov.tr/covid19api?getir=liste | cut -d "," -f5 | cut -d '"' -f4)

printf "Created by Burak Can\nGithub: github.com/Shootzz\n\n"
printf -v ay "%(%B)T"

printf " Vaka sayısı     : $vaka\n Tarih           : $gun $ay\n Test sayısı     : $test_sayi\n İyileşen sayısı : $iyilesen\n"

cat vakalar.txt &> /dev/null
if [ $? -ne 0 ] ; then
        printf "\n" > vakalar.txt
fi

grep "$gun $ay" vakalar.txt &> /dev/null
if [ $? -ne 0 ] ; then
        sed -i "1i[-] $gun $ay    : $vaka\n    Test sayısı     : $test_sayi\n    İyileşen sayısı : $iyilesen\n" vakalar.txt
else
        exit 1
fi
