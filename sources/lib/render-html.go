

package zscratchpad


import "html"




func DocumentRenderToHtml (_document *Document) (string, *Error) {
	
	if _document.RenderHtml != "" {
		return _document.RenderHtml, nil
	}
	
	_format := _document.Format
	if _format == "" {
		_format = "text"
		// return "", errorf (0xaff80238, "format empty")
	}
	
	_render := ""
	_error := (*Error) (nil)
	
	switch _format {
		
		case "text" :
			_render, _error = documentRenderTextToHtml (_document.BodyLines)
		
		case "snippets" :
			_render, _error = documentRenderSnippetsToHtml (_document.BodyLines)
		
		case "commonmark" :
			_render, _error = documentRenderCommonmarkToHtml (_document.BodyLines)
		
		case "gemini" :
			_render, _error = documentRenderGeminiToHtml (_document.BodyLines)
		
		default :
			return "", errorf (0xaf60ea6d, "format invalid `%s`", _document.Format)
	}
	
	if _error != nil {
		return "", _error
	}
	
	_render, _outcome, _error := DocumentSanitizeHtml (_document, _render)
	if _error != nil {
		return "", _error
	}
	
	_document.RenderHtml = _render
	_document.HtmlLinks = _outcome.UrlsLabel
	
	return _document.RenderHtml, nil
}




func documentRenderCommonmarkToHtml (_source []string) (string, *Error) {
	return parseAndRenderCommonmarkToHtml (_source)
}

func documentRenderGeminiToHtml (_source []string) (string, *Error) {
	return parseAndRenderGeminiToHtml (_source)
}

func documentRenderSnippetsToHtml (_source []string) (string, *Error) {
	return parseAndRenderSnippetsToHtml (_source)
}




func documentRenderTextToHtml (_source []string) (string, *Error) {
	
	_buffer := BytesBufferNewSize (128 * 1024)
	defer BytesBufferRelease (_buffer)
	
	_buffer.WriteString ("<pre><code>")
	for _, _line := range _source {
		_line = html.EscapeString (_line)
		_buffer.WriteString (_line)
		_buffer.WriteString ("\n")
	}
	_buffer.WriteString ("</code></pre>\n")
	
	_output := string (_buffer.Bytes ())
	
	return _output, nil
}

