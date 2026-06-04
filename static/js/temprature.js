function convert() {
    let value = parseFloat(document.getElementById('value').value);
    const from = document.getElementById('from').value;
    const to = document.getElementById('to').value;
    const resultElement = document.getElementById('result-value');

    if (!value || isNaN(value)) {
        resultElement.textContent = "0";
        return;
    }

    let result = value;

    // Convert everything to Celsius first, then to target unit
    let celsius = value;

    switch (from) {
        case "c":
            celsius = value;
            break;
        case "f":
            celsius = (value - 32) * 5 / 9;
            break;
        case "k":
            celsius = value - 273.15;
            break;
    }

    // Convert Celsius to target unit
    switch (to) {
        case "c":
            result = celsius;
            break;
        case "f":
            result = (celsius * 9 / 5) + 32;
            break;
        case "k":
            result = celsius + 273.15;
            break;
    }

    // Show result with 4 decimal places
    resultElement.textContent = result.toFixed(4);
}