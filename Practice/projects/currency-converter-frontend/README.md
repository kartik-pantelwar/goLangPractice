# Currency Converter Frontend

A modern, responsive frontend for your Go currency converter backend - **UI only, ready for your backend implementation**.

## Features

- 🎨 Modern, responsive design
- 💱 Currency swap functionality
- 📱 Mobile-friendly interface
- ⚡ Fast and intuitive UX
- 🛠️ Ready for backend integration

## Setup

1. **Open the frontend:**
   - Simply open `index.html` in your browser
   - Or use a local server like Live Server extension in VS Code

2. **Backend Integration:**
   - The frontend expects your Go backend at `http://localhost:8080/api`
   - Currently shows placeholder error until you implement your backend

## Backend Implementation Needed

You need to implement these endpoints in your Go backend:

### 1. Get All Currencies
```
GET /api/currencies
```
**Expected Response:**
```json
{
  "usd": "US Dollar",
  "eur": "Euro",
  "gbp": "British Pound Sterling",
  "jpy": "Japanese Yen",
  ...
}
```

### 2. Convert Currency
```
GET /api/convert?from=usd&to=eur&amount=100
```
**Expected Response:**
```json
{
  "from": "usd",
  "to": "eur",
  "amount": 100,
  "result": 85.23,
  "rate": 0.8523
}
```

## Current State

- ✅ **UI Components:** All HTML/CSS completed
- ✅ **Form Handling:** Basic validation and UI interactions
- ✅ **Currency Swap:** Button functionality implemented  
- ✅ **Loading States:** Spinner and error display ready
- ❌ **API Calls:** Removed - ready for your Go implementation
- ❌ **Currency Loading:** Uses placeholder data until backend ready

## TODO: Your Go Backend Tasks

1. **Create HTTP Server:**
   ```go
   // Example structure
   http.HandleFunc("/api/currencies", getCurrenciesHandler)
   http.HandleFunc("/api/convert", convertHandler)
   http.ListenAndServe(":8080", nil)
   ```

2. **Implement Currency List Endpoint:**
   - Fetch from exchange API: `https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies.json`
   - Return as JSON

3. **Implement Conversion Endpoint:**
   - Fetch rates from: `https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/{from}.json`
   - Calculate conversion
   - Return result with rate

4. **Add CORS Headers:**
   ```go
   w.Header().Set("Access-Control-Allow-Origin", "*")
   w.Header().Set("Access-Control-Allow-Methods", "GET")
   ```

## File Structure

```
currency-converter-frontend/
├── index.html      # Complete HTML structure
├── styles.css      # Complete CSS styling  
├── script.js       # UI logic + TODO backend calls
└── README.md       # This file
```

## JavaScript Integration Points

In `script.js`, you need to implement:

1. **`loadCurrencies()` method:**
   - Replace placeholder currencies with API call to your backend

2. **`convertCurrency()` method:**
   - Replace the TODO comment with actual fetch call to your backend

Example implementation:
```javascript
// In loadCurrencies method:
fetch(`${this.apiBaseUrl}/currencies`)
    .then(response => response.json())
    .then(currencies => this.populateCurrencySelects(currencies))
    .catch(error => console.error('Error:', error));

// In convertCurrency method:  
fetch(`${this.apiBaseUrl}/convert?from=${fromCurrency}&to=${toCurrency}&amount=${amount}`)
    .then(response => response.json())
    .then(result => this.showResult(result, amount, fromCurrency, toCurrency))
    .catch(error => this.showError('Conversion failed. Please try again.'));
```

## Testing

- **Frontend Only:** Currently shows "Backend not implemented yet" message
- **With Backend:** Will work seamlessly once your Go server is running

Happy coding with Go! 🚀
