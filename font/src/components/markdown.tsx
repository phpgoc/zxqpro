import React, { useState, useEffect } from 'react';
import markdownIt from 'markdown-it';
import hljs from 'highlight.js';
import 'highlight.js/styles/monokai.css'
import markdownItHighlightjs from 'markdown-it-highlightjs';
import { localeLanguage, serverUrl } from "../services/utils.ts";

interface MarkdownViewerProps {
  filePath: string; // Markdown 文件地址（相对路径或绝对路径）
}

const MarkdownViewer: React.FC<MarkdownViewerProps> = ({ filePath }) => {
  const [htmlContent, setHtmlContent] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  // 初始化 Markdown 解析器（可配置插件）
  const md: markdownIt = markdownIt({
    html: true,        // 允许 HTML 标签
    linkify: true,     // 自动转换链接
    typographer: true  // 转换特殊字符（如智能引号）
  }).use(markdownItHighlightjs, {
    hljs,
    auto: true,
    code: true
  });

  // 加载并解析 Markdown 文件
  useEffect(() => {
    const fetchMarkdown = async (): Promise<void> => {
      try {
        const filePathWithServerUrl: string = serverUrl() + "static/markdown/" + localeLanguage() + "/" + filePath;
        const response: Response = await fetch(filePathWithServerUrl);
        if (!response.ok) setError(`HTTP error! status: ${response.status}`);

        const text: string = await response.text();
        setHtmlContent(md.render(text)); // 直接转换为 HTML
        setIsLoading(false);
        setError(null)
      } catch (err: unknown) {
        if (err instanceof Error) {
          setError(`加载失败：${err.message}`);
        } else {
          setError(`加载失败：未知错误`);
        }
        setIsLoading(false);
      }
    };

    fetchMarkdown().then();
  }, [filePath]); // 监听文件路径变化

  return (
    <div className="markdown-viewer">
      {isLoading && <div className="loading">加载中...</div>}
      {error && <div className="error">{error}</div>}

      {/* 渲染解析后的 HTML（注意安全：确保文件来源可信） */}
      <div
        className="markdown-content"
        dangerouslySetInnerHTML={{ __html: htmlContent }}
      />
    </div>
  );
};

export default MarkdownViewer;
