# Use slim Python image
FROM python:3.11-slim

# Set working directory
WORKDIR /app

# Install dependencies first (layer caching)
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy application files
COPY server.py .
COPY index.html .
COPY style.css .
COPY app.js .

# Expose port
EXPOSE 5000

# Run the Flask app
CMD ["python", "server.py"]
