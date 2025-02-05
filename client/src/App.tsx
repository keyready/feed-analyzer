import { Sidebar } from './components/Sidebar';
import { MainPage } from './pages/MainPage';

export const App = () => {
    return (
        <section className="flex max-h-screen w-full overflow-hidden">
            <Sidebar />
            <MainPage />
        </section>
    );
};
