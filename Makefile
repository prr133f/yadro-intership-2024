.SILENT: build-image run

build-image:
	docker build -t yadro-intership-2024 .

run:
	docker run -it --rm yadro-intership-2024 data.txt