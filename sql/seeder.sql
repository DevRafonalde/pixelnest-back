\connect pixelnest

DO
$$
BEGIN
    -- Inserir valores na tabela t_perfis apenas se eles ainda não existirem
    IF NOT EXISTS (SELECT 1 FROM t_perfis WHERE nome = 'Administrador de Sistema') THEN
        INSERT INTO t_perfis (nome, descricao, ativo, data_ultima_atualizacao)
        VALUES ('Administrador de Sistema', 'Perfil de administrador com todas as permissões', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_perfis WHERE nome = 'Administrador de Operação') THEN
        INSERT INTO t_perfis (nome, descricao, ativo, data_ultima_atualizacao)
        VALUES ('Administrador de Operação', 'Perfil de administrador com todas as permissões da parte de operação', TRUE, NOW());
    END IF;

    -- Inserir valores na tabela t_permissoes apenas se eles ainda não existirem

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'delete-cliente-by-id') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, data_ultima_atualizacao)
        VALUES ('delete-cliente-by-id', 'Permissão para deletar um cliente', TRUE, NOW());
    END IF;

    -- Inserir valores na tabela t_parametros apenas se eles ainda não existirem
    
END
$$;
