interface MainPageProps {
    className?: string;
}

export const MainPage = (props: MainPageProps) => {
    return (
        <section className="p-10">
            <h1 className="text-3xl">
                Узнать, насколько ты пригоден к производству хохлятского сала или написанию
                бесполезных строк кода можно здесь.
            </h1>
            <h2 className="text-2xl italic">
                Все очень удобно, <u>уникальный</u> проект, разработанный <u>отечественной</u>{' '}
                компанией
            </h2>
        </section>
    );
};
