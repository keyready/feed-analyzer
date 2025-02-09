import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { HeroUIProvider } from '@heroui/react';

import { App } from './App.tsx';

import './styles/index.scss';
import './styles/tailwind.css';

createRoot(document.getElementById('root')!).render(
    <HeroUIProvider>
        <StrictMode>
            <App />
        </StrictMode>
    </HeroUIProvider>,
);
