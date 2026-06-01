
        function convert() {
            let value = parseFloat(document.getElementById('value').value);
            const from = document.getElementById('from').value;
            const to = document.getElementById('to').value;

            if (!value || isNaN(value)) {
                document.getElementById('result-value').textContent = "0";
                return;
            }

            // Convert to Gram (base unit)
            let grams = value;
            switch(from) {
                case "mg": grams = value / 1000; break;
                case "g": grams = value; break;
                case "kg": grams = value * 1000; break;
                case "ton": grams = value * 1000000; break;
                case "lb": grams = value * 453.59237; break;
                case "oz": grams = value * 28.349523125; break;
            }

            // Convert Gram to target unit
            let result = grams;
            switch(to) {
                case "mg": result = grams * 1000; break;
                case "g": result = grams; break;
                case "kg": result = grams / 1000; break;
                case "ton": result = grams / 1000000; break;
                case "lb": result = grams / 453.59237; break;
                case "oz": result = grams / 28.349523125; break;
            }

            document.getElementById('result-value').textContent = result.toFixed(4);
        }
   