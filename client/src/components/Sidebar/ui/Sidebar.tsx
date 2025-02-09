import { useState } from 'react';

interface SidebarProps {
    className?: string;
}

export const Sidebar = (props: SidebarProps) => {
    const [expanded, setExpanded] = useState<boolean>(true);

    return (
        <aside
            className={`flex h-screen flex-col justify-between bg-red-700 duration-200 ${expanded ? 'w-64' : 'w-14'}`}
        >
            <h1 className="text-center text-3xl">Оценка ценности</h1>

            <section className="flex flex-col">
                <a className="text-xl duration-200 hover:text-green-500" href="/">
                    Лидерборд
                </a>
                <a className="text-xl duration-200 hover:text-green-500" href="/">
                    Пройти тестирование
                </a>
            </section>

            <button
                className="cursor-pointer bg-green-800 text-white duration-200 hover:bg-green-500"
                onClick={() => setExpanded((ps) => !ps)}
            >
                {expanded ? 'Скрыть' : 'Открыть'}
            </button>
        </aside>
    );
};
