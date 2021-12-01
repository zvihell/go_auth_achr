import React from 'react';

const Home = (props: { name: string }) => {
    return (
        <div>
            {props.name ? 'Приветствую ' + props.name : 'Вы не авторизованы!'}
        </div>
    );
};

export default Home;
