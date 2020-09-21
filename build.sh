cd mains/maei
go build -o maei
cd ../..

cd mains/pan
go build -o pan
cd ../..

cd mains/wang
go build -o wang
cd ../..

cd mains/raei
go build -o raei
cd ../..

mv mains/maei/maei mains/pan/pan mains/wang/wang mains/raei/raei .