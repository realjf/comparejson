# ##############################################################################
# # File: Makefile                                                             #
# # Project: comparejson                                                       #
# # Created Date: 2023/12/05 10:57:17                                          #
# # Author: realjf                                                             #
# # -----                                                                      #
# # Last Modified: 2023/12/05 10:57:34                                         #
# # Modified By: realjf                                                        #
# # -----                                                                      #
# #                                                                            #
# ##############################################################################


B ?= master
M ?= "update"

.PHONY: push
push:
	@git add -A && git commit -m ${M} && git push origin ${B}

