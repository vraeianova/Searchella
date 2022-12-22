import base64
import json
import requests

user = "admin"
password = "Complexpass#123"
bas64encoded_creds = base64.b64encode(bytes(user + ":" + password, "utf-8")).decode("utf-8")

params = {
    "search_type": "match",
    "query":
    {
        "term": "a",
        
    },
    "from": 0, # use together with max_results for paginated results.
    "max_results": 20,
    "_source": [] # Leave this as empty array to return all fields.
}

# params = {
#     "search_type": "querystring",
#     "query":
#     {
#         "term": "+City:Turin +Silver",
#         "start_time": "2021-06-02T14:28:31.894Z",
#         "end_time": "2021-12-02T15:28:31.894Z"
#     },
#     "_source": ["_all"]
# }

headers = {"Content-type": "application/json",  "Authorization": "Basic " + bas64encoded_creds}
index = "testing"
zinc_host = "http://localhost:4080"
zinc_url = zinc_host + "/api/" + index + "/_search"

res = requests.post(zinc_url, headers=headers, data=json.dumps(params))

print(res.text)
