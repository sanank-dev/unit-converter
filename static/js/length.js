 function convert() {
            const value = parseFloat(document.getElementById('value').value);
            const from = document.getElementById('from').value;
            const to = document.getElementById('to').value;

            if (!value || isNaN(value)) {
                document.getElementById('result-value').textContent = "0";
                return;
            }

            // Convert to meter first
            let meters = value;
            switch(from) {
                case "mm": meters = value / 1000; break;
                case "cm": meters = value / 100; break;
                case "m": meters = value; break;
                case "km": meters = value * 1000; break;
                case "inch": meters = value * 0.0254; break;
                case "foot": meters = value * 0.3048; break;
                case "yard": meters = value * 0.9144; break;
                case "mile": meters = value * 1609.34; break;
            }

            // Convert meter to target unit
            let result = meters;
            switch(to) {
                case "mm": result = meters * 1000; break;
                case "cm": result = meters * 100; break;
                case "m": result = meters; break;
                case "km": result = meters / 1000; break;
                case "inch": result = meters / 0.0254; break;
                case "foot": result = meters / 0.3048; break;
                case "yard": result = meters / 0.9144; break;
                case "mile": result = meters / 1609.34; break;
            }

            document.getElementById('result-value').textContent = result.toFixed(4);
        }