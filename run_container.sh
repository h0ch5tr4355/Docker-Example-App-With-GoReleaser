container_name=Go_Releaser_Sample

if docker ps --format "{{.Names}}" --all | grep -q ${container_name}; then
    docker rm -f ${container_name}
    echo deleted
fi

docker run -d -p 8081:8080 -p 4000:4000  --name ${container_name} goreleaser_sample
echo created