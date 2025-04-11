// UserContext.tsx
import React, { createContext, useContext, useState, useEffect } from 'react';
import { UserInfo } from "../types/response.ts";

// 假设 UserInfo 类型定义如下


const UserContext = createContext<{
  user: UserInfo;
  updateUser: (newUser: UserInfo) => void;
} | null>(null);

const UserProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<UserInfo>(() => {
    const storedUser = localStorage.getItem("userInfo");
    return storedUser ? JSON.parse(storedUser) as UserInfo : {} as UserInfo;
  });

  const updateUser = (newUser: UserInfo) => {
    setUser(newUser);
    localStorage.setItem("userInfo", JSON.stringify(newUser));
  };

  // 监听 localStorage 变化
  useEffect(() => {
    const handleStorageChange = () => {
      const storedUser = localStorage.getItem("userInfo");
      if (storedUser) {
        setUser(JSON.parse(storedUser) as UserInfo);
      }
    };

    window.addEventListener('storage', handleStorageChange);

    return () => {
      window.removeEventListener('storage', handleStorageChange);
    };
  }, []);

  return (
    <UserContext.Provider value={{ user, updateUser }}>
      {children}
    </UserContext.Provider>
  );
};

const useUserContext = () => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error('useUserContext must be used within a UserProvider');
  }
  return context;
};

export { UserProvider, useUserContext };