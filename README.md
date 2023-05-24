### Generate mattermost webhook

    follow instructions [hear](https://developers.mattermost.com/integrate/webhooks/incoming/) for generate webhook

### set environment variables

    ```
    export MATTERMOSTWEBHOOK='your webhook endpoint'

    # link should be fleetdm.com/software/manage
    export TITLELINK='link'
    ```

### Run flogo app

    follow instructions [hear](https://github.com/Improwised/isight/issues/25) for run flogo app
    this will generate api `POST hostname:port/vulnerability`

### api 
    Request body

    
    {
    "timestamp": "0000-00-00T00:00:00Z",
    "vulnerability": {
    "cve": "CVE-2014-9471",
    "details_link": "https://nvd.nist.gov/vuln/detail/CVE-2014-9471",
    "hosts_affected": [
              {
                "id": 1,
                "display_name": "macbook-1",
                "url": "https://fleet.example.com/hosts/1"
                },
                {
                "id": 2,
                "display_name": "macbook-2",
                "url": "https://fleet.example.com/hosts/2"
                }
            ]
        }
    }

     
    response
    ![image](https://github.com/chintansakhiya/isight/assets/123356373/0f5e7599-39d7-4810-b3fb-7ba2cdbadb61)
    

### Generate fleet-dm webhook

    Select software > Manage automations
    Under Manage automations select webhook.
    Under Destination URL past flogo generated api

 
