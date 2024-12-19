// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ClassStorage, ClassDataStorage} from "../types/Class.sol";

library ClassUtils {
    function copyToStorage(
        ClassStorage memory m,
        ClassStorage storage $
    ) internal {
        $.id = m.id;
        $.name = m.name;
        $.symbol = m.symbol;
        $.description = m.description;
        $.uri = m.uri;
        $.uri_hash = m.uri_hash;
        $.data = ClassDataStorage({
            metadata: m.data.metadata,
            parent: m.data.parent,
            config: m.data.config,
            blind_box_state: m.data.blind_box_state
        });
    }
}
