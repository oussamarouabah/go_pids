path="Instances/graph_dolphins.txt"

echo "Run raei on $path"
./raei $path
echo "Run maei on $path"
./maei $path
echo "Run wang on $path"
./wang $path
echo "Run pan on $path"
./pan $path