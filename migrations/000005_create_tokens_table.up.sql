create table if not exists tokens(
    
    --sha256 hash of token is stored in tokens
    hash bytea primary key,
    
    --hace que se borren las entradas de aqui si se borra el user de alla
    --on delete restrict
    user_id bigint not null references users on delete cascade,

    --this will be 3 days 
    expiry timestamp(0) with time zone not null,

    --activation or authentication
    scope text not null
);