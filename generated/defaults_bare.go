package generated

// Do not edit, this file is automatically generated.

// Defaults_bare: scaffolding used in 'build' command
var Defaults_bare = map[string][]byte{
	"/.gitignore": []byte(`public
node_modules`),
	"/assets/favicon.svg": []byte(`<?xml version="1.0" encoding="UTF-8"?>
<!-- Created with Inkscape (http://www.inkscape.org/) -->
<svg width="47.596mm" height="47.596mm" version="1.1" viewBox="0 0 47.596 47.596" xmlns="http://www.w3.org/2000/svg" xmlns:cc="http://creativecommons.org/ns#" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">
 <metadata>
  <rdf:RDF>
   <cc:Work rdf:about="">
    <dc:format>image/svg+xml</dc:format>
    <dc:type rdf:resource="http://purl.org/dc/dcmitype/StillImage"/>
    <dc:title/>
   </cc:Work>
  </rdf:RDF>
 </metadata>
 <g transform="translate(-76.94 -96.762)" fill-rule="evenodd">
  <rect x="78.35" y="98.172" width="44.777" height="44.777" rx="4.5237" ry="4.5237" fill="#22a6ed" stroke="#22a6ed" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.8191"/>
  <path d="m88.5 112.76v26.379h8.4959v-15.879z" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.2401"/>
  <ellipse cx="101.01" cy="113.8" rx="13.633" ry="12.244" fill="#fff"/>
  <circle cx="97.99" cy="116.82" r="3.0232" fill="#22a6ed"/>
 </g>
</svg>
`),
	"/content/index.json": []byte(`{
	"title": "My Plenti Site"
}`),
	"/layout/content/index.svelte": []byte(`<script>
	export let title;
</script>

<h1>{title}</h1>

<style>
	h1 {
		font-family: Helvetica, sans-serif;
	}
</style>`),
	"/layout/global/head.svelte": []byte(`<script>
  export let title;
</script>

<head>
  <meta charset='utf-8'>
  <meta name='viewport' content='width=device-width,initial-scale=1'>

  <title>{ title }</title>

  <script type="module" src="/spa/ejected/main.js"></script>

  <link rel="icon" type="image/svg+xml" href="/assets/favicon.svg">
  <link rel='stylesheet' href='/spa/bundle.css'>
</head>
`),
	"/layout/global/html.svelte": []byte(`<script>
  import Head from './head.svelte';

  export let route, content, allContent;
</script>

<html lang="en">
<Head title={content.filename} />
<body>
  <svelte:component this={route} {...content.fields} {allContent} />
</body>
</html>`),
	"/package.json": []byte(`{
  "name": "my-plenti-app",
  "version": "1.0.0",
  "type": "module",
  "private": true,
  "dependencies": {
    "navaid": "^1.2.0",
    "regexparam": "^1.3.0",
    "svelte": "^3.29.4"
  }
}
`),
	"/plenti.json": []byte(`{
	"build": "public",
	"local": {
		"port": 3000
	}
}`),
}
