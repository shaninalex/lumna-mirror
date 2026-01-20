@pid=$$(sudo lsof -t -i :8000); \
if [ -n "$$pid" ]; then \
	echo "Killing process on port 8000 (PID: $$pid)"; \
	sudo kill -9 $$pid; \
else \
	echo "No process is listening on port 8000."; \
fi
