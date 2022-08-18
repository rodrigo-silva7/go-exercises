create table if not exists personalidades (
   id integer not null auto_increment,
   nome varchar(256),
   historia varchar(256),
   primary key(id)
);

insert into personalidades(nome,historia) values (
   ("Dalcídio Jurandir","Nascido em Cachoeira do Arari.... foi pra URSS bla bla"),
   ("Inglês de Sousa", "Escritor paraense nascido em....")
);
