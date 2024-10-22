Binance Pay Transactions API Example (Go)
=========================================

This Go program interacts with the Binance API to retrieve a list of Pay transactions. It sends a request to the ``/sapi/v1/pay/transactions`` endpoint, parses the response, and extracts relevant information like ``note``, ``amount``, ``currency``, and ``transactionTime`` from each transaction.

Features
--------

- Retrieves Binance Pay transactions using the ``/sapi/v1/pay/transactions`` endpoint.
- Automatically generates the required HMAC SHA256 signature for authenticated API requests.
- Parses and displays transaction details from the Binance response in the terminal.

Prerequisites
-------------

1. Go installed on your system. You can download and install Go from `here <https://golang.org/dl/>`_.
2. Binance API key and secret. You can generate these from your `Binance account <https://www.binance.com/en/my/settings/api-management>`_.

Setup
-----

1. Clone the repository or copy the code.

2. Replace the placeholders for ``apiKey`` and ``apiSecret`` in the code with your actual Binance API credentials.

.. code-block:: go

   const (
       apiKey    = "YOUR_API_KEY"
       apiSecret = "YOUR_API_SECRET"
       baseURL   = "https://api.binance.com"
   )

3. Install Go packages if required:
   - The program uses standard Go packages (``crypto/hmac``, ``crypto/sha256``, ``encoding/hex``, ``encoding/json``, ``fmt``, ``io``, ``net/http``, and ``net/url``).

4. Build and run the program.

.. code-block:: bash

   go run main.go

Code Overview
-------------

Signature Generation
~~~~~~~~~~~~~~~~~~~~
The Binance API requires an HMAC SHA256 signature for authentication. The signature is generated using the ``createSignature()`` function, which takes the query string as input and returns the corresponding signature.

.. code-block:: go

   func createSignature(data string) string {
       h := hmac.New(sha256.New, []byte(apiSecret))
       h.Write([]byte(data))
       return hex.EncodeToString(h.Sum(nil))
   }

API Request
~~~~~~~~~~~
The request is sent to the Binance ``/sapi/v1/pay/transactions`` endpoint using the HTTP ``GET`` method. The query parameters include a ``timestamp`` and the signature for authentication.

.. code-block:: go

   endpoint := "/sapi/v1/pay/transactions"
   timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())

   query := url.Values{}
   query.Set("timestamp", timestamp)
   signature := createSignature(query.Encode())
   query.Set("signature", signature)
   fullURL := fmt.Sprintf("%s%s?%s", baseURL, endpoint, query.Encode())

   req, err := http.NewRequest("GET", fullURL, nil)

Response Parsing
~~~~~~~~~~~~~~~~
The response is parsed into a custom ``Response`` struct, and the transactions are iterated over to extract and print the relevant details.

.. code-block:: go

   for _, tx := range response.Data {
       fmt.Printf("Note: %s\n", tx.Note)
       fmt.Printf("Amount: %s\n", tx.Amount)
       fmt.Printf("Currency: %s\n", tx.Currency)
       fmt.Printf("Transaction Time: %d\n\n", tx.TransactionTime)
   }

Example Output
--------------

If the request is successful, the response will look something like this in your terminal:

.. code-block:: plaintext

   Note: 
   Amount: 10
   Currency: USDT
   Transaction Time: 1728633026317

   Note: Binance Pay
   Amount: 0.3
   Currency: STRAX
   Transaction Time: 1726075803997

License
-------

This project is licensed under the MIT License - see the ``LICENSE`` file for details.
