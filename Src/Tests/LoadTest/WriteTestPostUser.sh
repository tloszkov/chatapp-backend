#hey -n 10000 -c 10 -m POST -H "Content-Type: application/json" -D write_user.json http://localhost:8090/api/v1/user/ | tee hey_output_postuser.txt
hey -n 10 -c 10 -m POST -H "Content-Type: application/json" -D write_user.json http://localhost:8090/api/v1/user/ | tee hey_output_postuser.txt
#-n 1000: The total number of requests to make.
#-c 10: The number of concurrent requests to make.
