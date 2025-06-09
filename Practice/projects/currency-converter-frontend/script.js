class CurrencyConverter {
    constructor() {
        this.apiBaseUrl = 'http://localhost:8080/api'; // Your Go backend URL
        
        this.init();
    }

    init() {
        this.loadCurrencies();
        this.bindEvents();
    }

    bindEvents() {
        const form = document.getElementById('currencyForm');
        const swapBtn = document.getElementById('swapBtn');

        form.addEventListener('submit', (e) => {
            e.preventDefault();
            this.convertCurrency();
        });

        swapBtn.addEventListener('click', () => {
            this.swapCurrencies();
        });
    }

    loadCurrencies() {
        // TODO: Implement this method to load currencies from your Go backend
        // For now, we'll populate with some common currencies as placeholder
        const commonCurrencies = {
            'usd': 'US Dollar',
            'eur': 'Euro',
            'gbp': 'British Pound Sterling',
            'jpy': 'Japanese Yen',
            'cad': 'Canadian Dollar',
            'aud': 'Australian Dollar',
            'chf': 'Swiss Franc',
            'cny': 'Chinese Yuan',
            'inr': 'Indian Rupee'
        };
        
        this.populateCurrencySelects(commonCurrencies);
    }

    populateCurrencySelects(currencies) {
        const fromSelect = document.getElementById('fromCurrency');
        const toSelect = document.getElementById('toCurrency');
        
        // Clear existing options except the first one
        fromSelect.innerHTML = '<option value="">Select currency</option>';
        toSelect.innerHTML = '<option value="">Select currency</option>';

        // Add currencies to both selects
        Object.keys(currencies).forEach(code => {
            const option1 = new Option(`${code.toUpperCase()} - ${currencies[code]}`, code);
            const option2 = new Option(`${code.toUpperCase()} - ${currencies[code]}`, code);
            fromSelect.add(option1);
            toSelect.add(option2);
        });

        // Set default values
        fromSelect.value = 'usd';
        toSelect.value = 'eur';
    }

    convertCurrency() {
        const amount = document.getElementById('amount').value;
        const fromCurrency = document.getElementById('fromCurrency').value;
        const toCurrency = document.getElementById('toCurrency').value;

        if (!amount || !fromCurrency || !toCurrency) {
            this.showError('Please fill in all fields');
            return;
        }

        if (fromCurrency === toCurrency) {
            this.showError('Please select different currencies');
            return;
        }

        this.showLoading();

        // TODO: Implement your Go backend API call here
        // Example:
        // fetch(`${this.apiBaseUrl}/convert?from=${fromCurrency}&to=${toCurrency}&amount=${amount}`)
        //     .then(response => response.json())
        //     .then(result => this.showResult(result, amount, fromCurrency, toCurrency))
        //     .catch(error => this.showError('Conversion failed. Please try again.'));
        
        // Placeholder - remove this when you implement your backend
        setTimeout(() => {
            this.showError('Backend not implemented yet. Please implement your Go API.');
        }, 1000);
    }

    swapCurrencies() {
        const fromSelect = document.getElementById('fromCurrency');
        const toSelect = document.getElementById('toCurrency');
        
        const temp = fromSelect.value;
        fromSelect.value = toSelect.value;
        toSelect.value = temp;
    }

    showResult(result, originalAmount, fromCurrency, toCurrency) {
        const resultDiv = document.getElementById('result');
        const fromAmountSpan = document.getElementById('fromAmount');
        const toAmountSpan = document.getElementById('toAmount');
        const exchangeRateSpan = document.getElementById('exchangeRate');

        fromAmountSpan.textContent = `${originalAmount} ${fromCurrency.toUpperCase()}`;
        toAmountSpan.textContent = `${result.result} ${toCurrency.toUpperCase()}`;
        exchangeRateSpan.textContent = `1 ${fromCurrency.toUpperCase()} = ${result.rate} ${toCurrency.toUpperCase()}`;

        this.hideAll();
        resultDiv.classList.remove('hidden');
    }

    showLoading() {
        this.hideAll();
        document.getElementById('loading').classList.remove('hidden');
    }

    showError(message) {
        const errorDiv = document.getElementById('error');
        const errorMessage = document.getElementById('errorMessage');
        
        errorMessage.textContent = message;
        
        this.hideAll();
        errorDiv.classList.remove('hidden');
    }

    hideAll() {
        document.getElementById('result').classList.add('hidden');
        document.getElementById('loading').classList.add('hidden');
        document.getElementById('error').classList.add('hidden');
    }
}

// Initialize the converter when the page loads
document.addEventListener('DOMContentLoaded', () => {
    new CurrencyConverter();
});

// Add some utility functions for better UX
document.addEventListener('DOMContentLoaded', () => {
    // Auto-focus amount input
    document.getElementById('amount').focus();
    
    // Add enter key support for conversion
    document.addEventListener('keypress', (e) => {
        if (e.key === 'Enter' && document.activeElement.tagName !== 'BUTTON') {
            document.getElementById('currencyForm').dispatchEvent(new Event('submit'));
        }
    });
});
