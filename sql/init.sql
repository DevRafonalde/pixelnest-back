-- Cria o banco de dados se ele não existir
SELECT 'CREATE DATABASE pixelnest'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'pixelnest')\gexec

-- Conecta ao banco de dados pixelnest
\connect pixelnest;

-- Criação tabelas controle de acesso

-- Cria tabela de usuários se ela não existir
CREATE TABLE IF NOT EXISTS t_usuarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
	senha VARCHAR(255) NOT NULL,
    ativo BOOLEAN,
	token_reset_senha VARCHAR(255),
    data_ultima_atualizacao TIMESTAMP,
	senha_atualizada BOOLEAN
);

-- Cria tabela de perfis se ela não existir
CREATE TABLE IF NOT EXISTS t_perfis (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    ativo BOOLEAN,
    data_ultima_atualizacao TIMESTAMP
);

-- Cria tabela de permissões se ela não existir
CREATE TABLE IF NOT EXISTS t_permissoes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) UNIQUE NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    ativo BOOLEAN,
    data_ultima_atualizacao TIMESTAMP
);

-- Cria tabela de associação de usuários a perfis se ela não existir
CREATE TABLE IF NOT EXISTS t_usuario_perfil (
    id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL,
    perfil_id INT NOT NULL,
    data_hora TIMESTAMP,
    FOREIGN KEY (usuario_id) REFERENCES t_usuarios(id),
    FOREIGN KEY (perfil_id) REFERENCES t_perfis(id)
);

-- Cria tabela de associação de perfis a permissões se ela não existir
CREATE TABLE IF NOT EXISTS t_perfil_permissao (
    id SERIAL PRIMARY KEY,
    perfil_id INT NOT NULL,
    permissao_id INT NOT NULL,
    data_hora TIMESTAMP,
    FOREIGN KEY (perfil_id) REFERENCES t_perfis(id),
    FOREIGN KEY (permissao_id) REFERENCES t_permissoes(id)
);

-- Cria tabela de parâmetros
CREATE TABLE IF NOT EXISTS t_parametros (
	id SERIAL PRIMARY KEY,
	nome VARCHAR(30),
	descricao varchar(255),
	valor varchar(255)
);

-- Criação das tabelas principais

-- Tabela t_jogos
CREATE TABLE IF NOT EXISTS t_jogos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    sinopse VARCHAR(255),
    avaliacao NUMERIC(1,1), -- Avaliação entre 0 e 5
    genero VARCHAR(255)
);

-- Tabela t_produtos
CREATE TABLE IF NOT EXISTS t_produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    descricao VARCHAR(255),
    avaliacao INT, -- Avaliação entre 0 e 5
    genero VARCHAR(255)
);

-- Tabela t_avaliacoes
CREATE TABLE IF NOT EXISTS t_avaliacoes (
    id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL, -- Campo para chave estrangeira
    produto_id INT,          -- Campo para chave estrangeira
    jogo_id INT,             -- Campo para chave estrangeira
    nota INT, -- nota entre 0 e 5
    avaliacao VARCHAR(255)
);

-- Tabela t_favoritos
CREATE TABLE IF NOT EXISTS t_favoritos (
    id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL, -- Campo para chave estrangeira
    produto_id INT,          -- Campo para chave estrangeira
    jogo_id INT              -- Campo para chave estrangeira
);

-- Criação da tabela t_usuario_jogo
CREATE TABLE IF NOT EXISTS t_usuario_jogo (
    id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL,
    jogo_id INT NOT NULL,
    UNIQUE (usuario_id, jogo_id) -- Garante que um usuário não possa ser associado ao mesmo jogo mais de uma vez
);

-- Adicionando as chaves estrangeiras após a criação das tabelas
-- Chave estrangeira em t_avaliacoes para t_usuarios
ALTER TABLE t_avaliacoes
ADD CONSTRAINT fk_avaliacao_usuario
FOREIGN KEY (usuario_id) REFERENCES t_usuarios(id);

-- Chave estrangeira em t_avaliacoes para t_produtos
ALTER TABLE t_avaliacoes
ADD CONSTRAINT fk_avaliacao_produto
FOREIGN KEY (produto_id) REFERENCES t_produtos(id);

-- Chave estrangeira em t_avaliacoes para t_jogos
ALTER TABLE t_avaliacoes
ADD CONSTRAINT fk_avaliacao_jogo
FOREIGN KEY (jogo_id) REFERENCES t_jogos(id);

-- Chave estrangeira em t_favoritos para t_usuarios
ALTER TABLE t_favoritos
ADD CONSTRAINT fk_favoritos_usuario
FOREIGN KEY (usuario_id) REFERENCES t_usuarios(id);

-- Chave estrangeira em t_favoritos para t_produtos
ALTER TABLE t_favoritos
ADD CONSTRAINT fk_favoritos_produto
FOREIGN KEY (produto_id) REFERENCES t_produtos(id);

-- Chave estrangeira em t_favoritos para t_jogos
ALTER TABLE t_favoritos
ADD CONSTRAINT fk_favoritos_jogo
FOREIGN KEY (jogo_id) REFERENCES t_jogos(id);

-- Adicionando a chave estrangeira para o campo usuario_id
ALTER TABLE t_usuario_jogo
ADD CONSTRAINT fk_usuario
FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
ON DELETE CASCADE;

-- Adicionando a chave estrangeira para o campo jogo_id
ALTER TABLE t_usuario_jogo
ADD CONSTRAINT fk_jogo
FOREIGN KEY (jogo_id) REFERENCES t_jogos(id)
ON DELETE CASCADE;
