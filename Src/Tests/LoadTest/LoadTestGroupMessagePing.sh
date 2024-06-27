hey -n 10000 -c 10 http://localhost:8090/api/v1/groupmessage/ping | tee hey_output_groupmessage_ping.txt

#-n 1000: The total number of requests to make.
#-c 10: The number of concurrent requests to make.