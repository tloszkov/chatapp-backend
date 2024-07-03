#hey -n 10000 -c 10 -m POST -H "Content-Type: application/json" -D write_user.json http://localhost:8090/api/v1/user/ | tee hey_output_postuser.txt
hey -n 100 -c 10 -m POST -H "Content-Type: application/json" -D write_user.json http://localhost:8090/api/v1/user/

#-n 100: Number of requests to run (100 requests).
#-c 10: Number of concurrent workers to run (10 concurrent requests at a time).
#-m POST: HTTP method to use (POST in this case).
#-H "Content-Type: application/json": HTTP header to set the content type to JSON.
#-D create_user.json: File containing the JSON payload to send with the request.
#http://localhost:8090/api/v1/user: URL to send the POST requests to.