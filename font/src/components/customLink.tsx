import { useNavigate } from 'react-router-dom';
import React, { ReactNode, useState } from 'react';
import { localNavigate } from "../services/utils.ts";

export default function CustomLink({ to, children }: { to: string; children: ReactNode }) {
  const navigate = useNavigate();
  const [isContextMenuOpen, setIsContextMenuOpen] = useState(false);
  const [contextMenuPosition, setContextMenuPosition] = useState({ x: 0, y: 0 });



    function handleClick  (e: React.MouseEvent<HTMLAnchorElement>) {
      e.preventDefault();
      console.log(to)
      localNavigate(navigate,to)
    }



  function handleContextMenu(e: React.MouseEvent<HTMLAnchorElement>) {
    e.preventDefault();
    setIsContextMenuOpen(true);
    console.log(isContextMenuOpen)
    console.log({ x: e.clientX, y: e.clientY })
    setContextMenuPosition({ x: e.clientX, y: e.clientY });
  }

  function handleOutsideClick() {
    if (isContextMenuOpen) {
      setIsContextMenuOpen(false);
    }
  }

  function openInNewTab() {
    window.open(to, '_blank');
    setIsContextMenuOpen(false);
  }

  return (
    <div onClick={handleOutsideClick}>
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
