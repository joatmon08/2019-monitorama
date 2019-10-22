slide:
	reveal-md slides.md --theme night -w

pages:
	npm run generate

print:
	reveal-md slides.md --theme night --print slides.pdf