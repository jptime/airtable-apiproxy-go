# airtable-apiproxy-go

Proof of concept.  

HTTP reverse proxy to forward requests to airtables api.

Fetches single table.

edit .env file with your airtable apikey and the path to the table.  

Default port is 9090.  Pass in a different port via flags.

Lot's of TODOs.  But main ones being various endpoints and passing in filterFormulas to forward on to airtable.