import { useNavigate } from 'react-router-dom';
import React, { ReactNode, useEffect,useState } from "react";
import { localNavigate } from "../services/utils.ts";

export default function CustomLink({ to, children }: { to: string; children: ReactNode }) {
  const navigate = useNavigate();
  const [isContextMenuOpen, setIsContextMenuOpen] = useState(false);
  const [contextMenuPosition, setContextMenuPosition] = useState({ x: 0, y: 0 });

    function handleClick  (e: React.MouseEvent<HTMLAnchorElement>) {
      e.preventDefault();
      localNavigate(navigate,to)
    }

  useEffect(() => {
    document.addEventListener('mousedown', handleOutsideClick);
    return () => {
      document.removeEventListener('mousedown', handleOutsideClick);
    };
  }, [isContextMenuOpen]);


  function handleContextMenu(e: React.MouseEvent<HTMLAnchorElement>) {
    e.preventDefault();
    setIsContextMenuOpen(true);
    setContextMenuPosition({ x: e.clientX, y: e.clientY });
  }


  function handleOutsideClick(e: MouseEvent) {
    const target = e.target as HTMLElement;
    const linkElement = document.querySelector('a.text-blue-500.underline.cursor-pointer');
    const menuElement = document.querySelector('div[style*="position: fixed"]');

    if (isContextMenuOpen && linkElement && menuElement) {
      if (!linkElement.contains(target) &&!menuElement.contains(target)) {
        setIsContextMenuOpen(false);
      }
    }
  }

  function openInNewTab() {
    window.open(to, '_blank');
    setIsContextMenuOpen(false);
  }

  return (
    <div onClick={()=>handleOutsideClick}>
      <a
        href={to}
        onClick={handleClick}
        onContextMenu={handleContextMenu}
        className="text-blue-500 underline cursor-pointer"
      >
        {children}
      </a>
      {isContextMenuOpen && (
        <div
          style={{
            position: 'fixed',
            left: contextMenuPosition.x,
            top: contextMenuPosition.y,
            backgroundColor: 'white',
            border: '1px solid #ccc',
            boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
            zIndex: 1000
          }}
        >
          <ul style={{ listStyleType: 'none', padding: 0, margin: 0 }}>
            <li style={{ padding: '8px 16px', cursor: 'pointer' }} onClick={openInNewTab}>
              open in new tab
            </li>
          </ul>
        </div>
      )}
    </div>
  );
}
