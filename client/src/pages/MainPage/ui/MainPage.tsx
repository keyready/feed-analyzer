import { FormEvent, useCallback, useMemo, useState } from 'react';
import { Slider } from '@heroui/react';

enum SKILLS_TYPES {
    TechnicalSkills = 'Технические навыки',
    SoftSkills = 'Софт-скиллы',
    DomainKnowledge = 'Знания предметной области',
    AcademicAchievements = 'Академические достижения',
    TeamMethodicalSkills = 'Командно-методические навыки',
}

/**
 * SKILLS_TYPES.AcademicAchievements
 * [тип мероприятия (хак, цтф, выставка(?))] - [название] - [место] - [грамота(файл)]
 */

const SKILLS_NAMES: Record<SKILLS_TYPES, string[]> = {
    [SKILLS_TYPES.TechnicalSkills]: ['Программирование', 'Работа с БД', 'Сети'],
    [SKILLS_TYPES.SoftSkills]: ['Коммуникабельность', 'Лидерство', 'Вежливость'],
    [SKILLS_TYPES.DomainKnowledge]: ['АПР', 'РО', 'МД', 'ОКР'],
    [SKILLS_TYPES.AcademicAchievements]: ['АПР'],
    [SKILLS_TYPES.TeamMethodicalSkills]: [
        'Строевая подготовка',
        'Огневая подготовка',
        'РХБ подготовка',
    ],
};

export const MainPage = () => {
    const [firstname, setFirstname] = useState<string>('');
    const [lastname, setLastname] = useState<string>('');
    const [rank, setRank] = useState<string>('');
    const [age, setAge] = useState<number>(0);

    const [formStep, setFormStep] = useState<number>(1);

    const handleFormSubmit = useCallback(
        (ev: FormEvent<HTMLFormElement>) => {
            ev.preventDefault();
            if (formStep === 1) {
                const newSP = new URLSearchParams({
                    firstname,
                    lastname,
                    rank,
                });
                setFormStep(2);
            }
        },
        [firstname, formStep, lastname, rank],
    );

    const renderForm = useMemo(() => {
        switch (formStep) {
            case 1: {
                return (
                    <form onSubmit={handleFormSubmit}>
                        <input
                            value={firstname}
                            onChange={(ev) => setFirstname(ev.target.value)}
                            type="text"
                            name={'firstname'}
                        />
                        <input
                            value={lastname}
                            onChange={(ev) => setLastname(ev.target.value)}
                            type="text"
                            name={'lastname'}
                        />
                        <input
                            value={age}
                            onChange={(ev) => setAge(~~ev.target.value)}
                            type="number"
                            name={'age'}
                        />
                        <input
                            value={rank}
                            onChange={(ev) => setRank(ev.target.value)}
                            type="text"
                            name={'rank'}
                        />
                        <input type="file" name={'image'} />
                        <button type={'submit'}>Продолжить</button>
                    </form>
                );
            }

            case 2: {
                return (
                    <form onSubmit={handleFormSubmit}>
                        {Object.values(SKILLS_TYPES).map((key) => (
                            <form className="mb-10">
                                <h1 className="text-red-400">{key}</h1>
                                {Object.values(SKILLS_NAMES[key]).map((k) => (
                                    <div>
                                        <p key={k}>{k}</p>
                                        <Slider
                                            className="max-w-md"
                                            defaultValue={0.4}
                                            label="Temperature"
                                            maxValue={1}
                                            minValue={0}
                                            step={0.01}
                                        />
                                    </div>
                                ))}
                            </form>
                        ))}
                        <button type={'submit'}>Продолжить</button>
                    </form>
                );
            }
        }
    }, [age, firstname, formStep, handleFormSubmit, lastname, rank]);

    return (
        <section className="p-10">
            <h1>Шаг {formStep}</h1>
            {renderForm}
        </section>
    );
};
