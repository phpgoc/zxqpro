```typescript
import React, { useState, useEffect } from 'react';
import markdownIt from 'markdown-it';
import hljs from 'highlight.js';
import 'highlight.js/styles/default.css';
import markdownItHighlightjs from 'markdown-it-highlightjs';
import { localeLanguage, serverUrl } from "../services/utils.ts";

export default MarkdownViewer;

```